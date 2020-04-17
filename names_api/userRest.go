package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	log.Println("respondWithJSON start")
	defer log.Println("respondWithJSON done")
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithGenericErr(w http.ResponseWriter, code int, err error) {
	log.Println("respondWithGenericErr start")
	defer log.Println("respondWithGenericErr done")
	var resp genericError
	resp.Code = code
	log.Print("Return status code ", resp.Code)
	resp.Message = err.Error()
	log.Print(resp)
	respondwithJSON(w, resp.Code, resp)
}

func dbGetUsers() ([]*User, error) {
	var users []*User
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := db.mongoDatabase.Collection(config.MongoCollection)
	filter := bson.M{"is_deleted": bson.M{"$ne": true}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal("dbGetUsers error on find: ", err)
	}
	defer cur.Close(ctx)

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result User
		err := cur.Decode(&result)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("Appending user: %v", result)
			users = append(users, &result)
		}
	}
	return users, err
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("getUsers start")
	defer log.Println("getUsers")
	users, err := dbGetUsers()
	if err != nil {
		log.Print("Did not get the users", err)
		respondWithGenericErr(w, http.StatusInternalServerError, err)
		return
	}

	log.Println(users)
	respondwithJSON(w, http.StatusOK, users)
}

func dbGetUser(userID string) (User, error) {
	var err error
	var user User

	log.Printf("Finding on %s", userID)

	collection := db.mongoDatabase.Collection(config.MongoCollection)
	filter := bsonFilter(userID)
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		fmt.Printf("Error was: '%s'", err.Error())
		if strings.Contains(err.Error(), "no documents in result") {
			log.Print("FindOne ", err)
		} else {
			log.Print("Error on FindOne: ", err)
		}
	} else {
		log.Print(user)
	}
	return user, err
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	user, err := dbGetUser(userID)

	if err != nil {

		if strings.Contains(err.Error(), "mongo: no documents in result") {
			log.Print("got soft err: ", err)
			respondWithGenericErr(w, http.StatusNotFound, err)
			return
		} else {
			log.Print("got err: ", err)
			respondWithGenericErr(w, http.StatusInsufficientStorage, err)
			return
		}
	}

	log.Print(user)
	respondwithJSON(w, http.StatusOK, user)
}

func dbDeleteUser(userID string) error {
	log.Printf("Deleting on %s", userID)

	collection := db.mongoDatabase.Collection(config.MongoCollection)

	filter := bsonFilter(userID)
	update := bson.D{
		{"$set", bson.D{
			primitive.E{"is_deleted", true},
		}},
	}
	// deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Print("Error on UpdateOne (is_deleted)", err)
	} else {
		log.Print(updateResult)
		log.Printf("Matched %v documents and modified %v documents", updateResult.MatchedCount, updateResult.ModifiedCount)
	}
	return err
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	err := dbDeleteUser(userID)

	if err != nil {
		log.Print("got err: ", err)
		var resp genericError
		resp.Message = err.Error()
		if strings.Contains(err.Error(), "no documents in result") {
			respondWithGenericErr(w, http.StatusNotFound, err)
			return
		} else {
			respondWithGenericErr(w, http.StatusInternalServerError, err)
			return
		}
	}
	respondwithJSON(w, http.StatusNoContent, nil)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var reqData User
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		log.Print("Error Decoding body", err)
		respondWithGenericErr(w, http.StatusBadRequest, err)
		return
	}
	log.Print("Deserialized into User: ", reqData)

	vMsg := reqData.Validate()
	if len(vMsg) != 0 {
		log.Print("Validation Failed on user: ", vMsg)
		respondWithGenericErr(w, http.StatusBadRequest, errors.New(strings.Join(vMsg, ", ")))
		return
	}

	collection := db.mongoDatabase.Collection(config.MongoCollection)

	if collection != nil {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		res, err := collection.InsertOne(ctx, reqData)
		log.Print("res ", res)
		log.Print("err", err)
		if err != nil {
			log.Print("Failed to insert reqdata for user")
			// w.WriteHeader(http.StatusInternalServerError)
			respondWithGenericErr(w, http.StatusInternalServerError, err)
			return
		} else {
			id := res.InsertedID
			idStr := fmt.Sprintf("%v", id)
			log.Print("created! ", idStr)

			respondwithJSON(w, http.StatusCreated, nil)
			return
		}
	}
}

func bsonFilter(userID string) bson.D {
	objID, _ := primitive.ObjectIDFromHex(userID)
	return bson.D{
		{"is_deleted", bson.M{"$ne": true}},
		{"$or",
			bson.A{
				bson.D{{"_id", objID}},
				bson.D{{"user_name", userID}},
				bson.D{{"email", userID}},
			},
		},
	}
}

func putUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	var putUserData User
	if err := json.NewDecoder(r.Body).Decode(&putUserData); err != nil {
		log.Print("Error Decoding body", err)
		respondWithGenericErr(w, http.StatusBadRequest, err)
		return
	}
	log.Print("Deserialized into User: ", putUserData)

	vMsg := putUserData.Validate()
	if len(vMsg) != 0 {
		log.Print("Validation Failed on user: ", vMsg)
		respondWithGenericErr(w, http.StatusBadRequest, errors.New(strings.Join(vMsg, ", ")))
		return
	}

	collection := db.mongoDatabase.Collection(config.MongoCollection)

	filter := bsonFilter(userID)
	_, err := collection.ReplaceOne(context.TODO(), filter, putUserData)
	if err != nil {
		log.Print("Problem updated userId: ", userID)
		respondWithGenericErr(w, http.StatusInternalServerError, err)
	} else {
		respondwithJSON(w, http.StatusNoContent, nil)
	}
}

func patchUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	log.Print("patchUser ", userID)
	defer log.Print("patchUser end for ", userID)

	user, err := dbGetUser(userID)
	if err != nil {
		log.Print("Error getting updated user: ", userID, err)
		respondWithGenericErr(w, http.StatusNotFound, err)
		return
	}

	var patches jsonpatch.Patch
	if err := json.NewDecoder(r.Body).Decode(&patches); err != nil {
		log.Print("Error extracting patches", err)
		respondwithJSON(w, http.StatusBadRequest, err)
		return
	}
	log.Print("Patches ", patches)

	original, err := json.Marshal(user)
	if err != nil {
		log.Print("Error marshalling original to json: ", userID, err)
		respondWithGenericErr(w, http.StatusBadRequest, err)
		return
	}

	modifiedJSON, err := patches.Apply(original)
	if err != nil {
		log.Print("Could not apply patches to original", err)
		respondWithGenericErr(w, http.StatusBadRequest, err)
		return
	}
	var modifiedUser User
	err = json.Unmarshal(modifiedJSON, &modifiedUser)
	if err != nil {
		log.Print("Could not unmarshal the patched json to a user ", err)
		respondWithGenericErr(w, http.StatusBadRequest, err)
		return
	}

	vMsg := modifiedUser.Validate()
	if len(vMsg) != 0 {
		log.Print("Validation Failed on user: ", vMsg)
		respondWithGenericErr(w, http.StatusBadRequest, errors.New(strings.Join(vMsg, ", ")))
		return
	}

	collection := db.mongoDatabase.Collection(config.MongoCollection)

	filter := bsonFilter(userID)
	replaceResult, err := collection.ReplaceOne(context.TODO(), filter, modifiedUser)
	if err != nil {
		log.Print("Could not ReplaceOne on the patched user", err)
		respondWithGenericErr(w, http.StatusInternalServerError, err)
		return
	}
	log.Print(replaceResult)
	respondwithJSON(w, http.StatusOK, modifiedUser)
}

type OuterData struct {
	Results []PersonStruct `json:"results"`
}
type NameStruct struct {
	First string `json:"first"`
	Last  string `json:"last"`
}
type PersonStruct struct {
	Gender string     `json:"gender"`
	Name   NameStruct `json:"name"`
}

func populateData(w http.ResponseWriter, r *http.Request) {
	log.Println("populateData start")
	defer log.Println("populateData end")
	log.Println("Getting names")

	//https://randomuser.me/api/
	response, err := http.Get("https://randomuser.me/api/?results=100")
	if err != nil {
		log.Printf("get error %v", err)
		respondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}
	data, _ := ioutil.ReadAll(response.Body)
	log.Println(string(data))

	var parsedData OuterData
	json.Unmarshal([]byte(data), &parsedData)
	log.Printf("len results: %v", len(parsedData.Results))

	collection := db.mongoDatabase.Collection(config.MongoCollection)
	if collection == nil {
		respondWithGenericErr(w, http.StatusInternalServerError, nil)
		return
	}
	for x, r := range parsedData.Results {
		log.Printf("%v, %v %v", x, r.Name.First, r.Name.Last)

		var user User
		user.Name = fmt.Sprintf("%v %v", r.Name.First, r.Name.Last)

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		res, err := collection.InsertOne(ctx, user)
		log.Print("res ", res)
		log.Print("err", err)
		if err != nil {
			log.Print("Failed to insert reqdata for user")
			// w.WriteHeader(http.StatusInternalServerError)
			respondWithGenericErr(w, http.StatusInternalServerError, err)
			return
		}
		id := res.InsertedID
		idStr := fmt.Sprintf("%v", id)
		log.Print("created! ", idStr)
	}

	respondwithJSON(w, http.StatusCreated, nil)
}
