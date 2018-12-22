package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetRootPath Handle the root path
func GetRootPath(w http.ResponseWriter, r *http.Request) {
	contents := `{ "response" : "Hi, Ed!"}`
	json.NewEncoder(w).Encode(contents)
}

// Server does what
func Server() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetRootPath).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

// func main() {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/", GetRootPath).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":8080", router))

// }
