package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

//r is a pointer to the http.Request interface
func getSomething(w http.ResponseWriter, r *http.Request) {
	//:= infers type in a function
	dataResults := DataResults {
		DataResult{Id: "1", Name: "Tom", Content:"something blah blah"},
	}

	fmt.Println("Endpoint Hit: getSomething", r.Host)

	if r.Host == "localhost:8080" {
		//w is the response writer
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    	w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(dataResults); err != nil {
        	panic(err)
    	}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    	w.WriteHeader(http.StatusBadRequest)
    	fmt.Println("Not Running from localhost:8080. Running from: ", r.Host)
	}
}