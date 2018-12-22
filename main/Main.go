package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EdSwArchitect/brooklyn-es/query"
	"github.com/gorilla/mux"
)

// GetRootPath Handle the root path
func GetRootPath(w http.ResponseWriter, r *http.Request) {
	contents := `{ "response" : "Hi, Ed!"}`
	json.NewEncoder(w).Encode(contents)
}

// GetEvents Get the event
func GetEvents(w http.ResponseWriter, r *http.Request) {

	// a map of values
	params := mux.Vars(r)

	fmt.Printf("Params: %v\n", params)

	fmt.Printf("val is: '%s'\n", params["val"])

	queryResults := query.GetBodyQueryByEventsResty("darkstar", 9200, "mybrooklyn", params["val"], 0, 50)

	objectResults, _ := query.ParseElasticJSONObject(queryResults)

	json.NewEncoder(w).Encode(objectResults.Hits.Hits)

}

// Server does what
func Server() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetRootPath).Methods("GET")
	router.HandleFunc("/getEvents/{val}", GetEvents).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func main() {
	args := os.Args

	fmt.Printf("Hi, Ed: %v\n", args)

	uri := args[1]

	fmt.Printf("The one arg is: '%s'\n", uri)

	// get the information about the cluster
	status, _ := query.GetElasticServerInfo(uri)

	fmt.Printf("Name: %s\nVersion: %s\nLucene Version: %s\nCluster name: %s\nTagline: %s\n", status.Name, status.Version.Number,
		status.Version.LuceneVersion, status.ClusterName, status.Tagline)

	// queryResuts := query.GetQueryResults("darkstar", 9200, "mybrooklyn")

	// fmt.Printf("Query results\n%s\n\n*********\n\n", queryResuts)

	// queryResuts = query.GetLimitedQueryResults("darkstar", 9200, "mybrooklyn", 7100, 50)

	// fmt.Printf("Query results\n%s\n", queryResuts)

	queryResults := query.GetBodyQueryByEventsResty("darkstar", 9200, "mybrooklyn", "Father's", 0, 50)

	fmt.Printf("\n\nFLAG FLAG FLAG\n\nQuery results\n%s\n==========\n\n", queryResults)

	objectResults, _ := query.ParseElasticJSONObject(queryResults)

	fmt.Printf("The length is: %v\n", len(objectResults.Hits.Hits))

	for _, daHit := range objectResults.Hits.Hits {
		fmt.Printf("%v", daHit)
		fmt.Println("----")
	}

	Server()

}
