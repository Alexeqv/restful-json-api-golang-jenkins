package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Message string `json:"Message"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "API REST - Exmaple GO",
		Message: "HELLO WORLD!",
	},
	{
		ID:          "2",
		Title:       "Docker",
		Message: "docker run",
	},
}

func service(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/service", service).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
