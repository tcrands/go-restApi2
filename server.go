package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    handleRequests()
}

func handleRequests() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/get", getSomething)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func getSomething(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here you can get some data")
	fmt.Println("Endpoint Hit: getSomething")
}

type DataResult struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Content string `json:"content"`
}

type DataResults []DataResult 