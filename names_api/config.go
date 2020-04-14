package main

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	MongoURI          string
	MongoUser         string
	MongoPassword     string
	MongoServer       string
	MongoPort         string
	MongoDB           string
	MongoCollection   string
	ServerEnvironment string
	ServicePort       string
}

var (
	config Config
)

func (c *Config) init() {
	c.MongoDB = os.Getenv("MONGO_DB")
	c.MongoUser = os.Getenv("MONGO_USER")
	c.MongoPassword = os.Getenv("MONGO_PASSWORD")
	c.MongoServer = os.Getenv("MONGO_SERVER")
	c.MongoPort = os.Getenv("MONGO_PORT")
	c.MongoCollection = os.Getenv("MONGO_COLLECTION")
	c.ServerEnvironment = os.Getenv("SERVER_ENVIRONMENT")
	c.ServicePort = os.Getenv("SERVICE_PORT")

	c.MongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?retryWrites=false", c.MongoUser, c.MongoPassword, c.MongoServer, c.MongoPort, c.MongoDB)
	log.Print(c.MongoURI)
}
