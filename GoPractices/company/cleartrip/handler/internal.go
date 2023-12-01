package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
}

// InitHttpClient init internal url for upcoming request.
func InitHttpClient() {
	router.HandleFunc(GetHelloURL, GetHelloHandler).Methods("GET")
	router.HandleFunc(PostHelloURL, PostHelloHandler).Methods("POST")
	http.Handle("/", router)

	fmt.Println("Server listening on :", 8080)
	if err := http.ListenAndServe(HostAndPort, nil); err != nil {
		fmt.Println(err)
	}
}

// underlyingHandler call GetUnderlyingPricesHandler return all underlying response.
func GetHelloHandler(response http.ResponseWriter, req *http.Request) {
	data := "hello, good morning"
	// Convert the data to JSON format
	responseJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(response, "error", http.StatusInternalServerError)
		return
	}

	// Set the content type to application/json
	response.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	if _, err := response.Write(responseJSON); err != nil {
		fmt.Println(err)
	}
}

// underlyingHandler call GetUnderlyingPricesHandler return all underlying response.
func PostHelloHandler(response http.ResponseWriter, req *http.Request) {
	input := Input{}

	json.NewDecoder(req.Body).Decode(&input)
	fmt.Println(input)
	data := "hello, good morning"
	// Convert the data to JSON format
	responseJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(response, "error", http.StatusInternalServerError)
		return
	}

	// Set the content type to application/json
	response.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	if _, err := response.Write(responseJSON); err != nil {
		fmt.Println(err)
	}
}
