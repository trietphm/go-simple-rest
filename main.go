package main

import (
	"go-simple-rest/route"
	"log"
	"net/http"
)

func main() {
	router := route.Router()

	log.Fatal(http.ListenAndServe(":3000", router))
}
