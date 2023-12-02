package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	router := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", router))
}
