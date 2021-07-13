package main

import (
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HTTP Status code and an error message
type genericError struct {
	// the status code of the request
	Code    int    `json:"code"`
	Message string `json:"message"`
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
	return messages
}
