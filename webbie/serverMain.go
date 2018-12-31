package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EdSwArchitect/brooklyn-es/query"
	"github.com/gorilla/mux"
)

// GetRootPath Handle the root path
func GetRootPath(w http.ResponseWriter, r *http.Request) {
	contents := `{ "response" : "Hi, Ed! Webbie Root path"}`
	json.NewEncoder(w).Encode(contents)
}

// GetEvents Get the event based on the day event
func GetEvents(w http.ResponseWriter, r *http.Request) {

	// a map of values
	params := mux.Vars(r)

	fmt.Printf("Params: %v\n", params)

	fmt.Printf("val is: '%s'\n", params["val"])

	queryResults := query.GetBodyQueryByEventsResty("darkstar", 9200, "mybrooklyn", params["val"], 0, 50)

	objectResults, _ := query.ParseElasticJSONObject(queryResults)

	json.NewEncoder(w).Encode(objectResults.Hits.Hits)

}

// GetWeather Get the event based on the weather
func GetWeather(w http.ResponseWriter, r *http.Request) {

	// a map of values
	params := mux.Vars(r)

	fmt.Printf("Params: %v\n", params)

	fmt.Printf("val is: '%s'\n", params["val"])

	queryResults := query.GetByWeather("darkstar", 9200, "mybrooklyn", params["val"], 0, 50)

	objectResults, _ := query.ParseElasticJSONObject(queryResults)

	json.NewEncoder(w).Encode(objectResults.Hits.Hits)

}

// Server sets up the RESTful server with the URIs given
func Server() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetRootPath).Methods("GET")
	router.HandleFunc("/getEvents/{val}", GetEvents).Methods("GET")
	router.HandleFunc("/getWeather/{val}", GetWeather).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
