package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func getSomething(w http.ResponseWriter, r *http.Request) {

	dataResults := DataResults {
		DataResult{Id: "1", Name: "Tom", Content:"something blah blah"},
	}

	fmt.Println("Endpoint Hit: getSomething")

	if err := json.NewEncoder(w).Encode(dataResults); err != nil {
        panic(err)
    }
}