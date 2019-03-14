package main

import (
	route "github.com/heaptracetechnology/microservice-telegram/route"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
