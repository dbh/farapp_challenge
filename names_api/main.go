package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

var (
	router chi.Router
)

func main() {
	config.init()

	err := db.init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.close()

	log.Println("Starting chi router")
	initRouter()
	setupUserRoutes()

	log.Println("Starting service on port ", config.ServicePort)
	http.ListenAndServe(":"+config.ServicePort, router)
	log.Println("Exiting")
}
