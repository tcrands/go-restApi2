package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func main() {
    handleRequests()
}

func handleRequests() {

	mainRouter := mainRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}


