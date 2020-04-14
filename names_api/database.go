package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct {
	mongoClient   *mongo.Client
	mongoDatabase *mongo.Database
}

var (
	db Db
)

func (db *Db) init() error {
	var err error
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	db.mongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = db.mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	db.mongoDatabase = db.mongoClient.Database(config.MongoDB)
	return err
}

func (db *Db) close() {
	db.mongoClient.Disconnect(context.TODO())
}
