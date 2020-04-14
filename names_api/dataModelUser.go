package main

import (
	"log"
	// "time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HTTP Status code and an error message
type genericError struct {
	// the status code of the request
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type swaggUserResp struct {
	Body struct {
		// HTTP status code 200 - Status OK
		Code int  `json:"code"`
		Data User `json:"user"`
	} `json:"body"`
}

type swaggUsersResp struct {
	Body struct {
		// HTTP status code 200 - Status OK
		Code int     `json:"code"`
		Data []*User `json:"users"`
	} `json:"body"`
}

type User struct {
	Id    *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string              `json:"name" bson:"name"`
	Email string              `json:"email,omitempty" bson:"email,omitempty"`
	Phone string              `json:"phone,omitempty" bson:"phone,omitempty"`
}

// add modified at?

func (u *User) Validate() []string {
	log.Print("Doing validation on: ", u)
	var messages []string
	if u.Name == "" {
		messages = append(messages, "Name is required")
	}
	// if u.Email == "" {
	// 	messages = append(messages, "Email is required")
	// }
	// if u.Phone == "" {
	// 	messages = append(messages, "Phone is required")
	// }
	return messages
}
