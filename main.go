package main

import (
	"go-simple-rest/resource"
	"go-simple-rest/route"
	"log"
	"net/http"
)

func main() {
	router := route.Router()
	err := resource.InitDB()
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":3000", router))
}
