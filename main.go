package main

import (
	"log"
	"net/http"
	_ "expvar"
)

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
