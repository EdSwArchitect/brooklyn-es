package main

import (
	sql "database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/EdSwArchitect/brooklyn-es/query"
	webbie "github.com/EdSwArchitect/brooklyn-es/webbie"
	"github.com/gorilla/mux"
)

// SimpleBrooklyn type
type SimpleBrooklyn struct {
	StartDate     string  `json:"startDate"`
	Weather       string  `json:"weather"`
	Temperature   float32 `json:"temperature"`
	Precipitation float32 `json:"precipitation"`
	Events        string  `json:"events"`
}

// GetRootPath Handle the root path
// func GetRootPath(w http.ResponseWriter, r *http.Request) {
// 	contents := `{ "response" : "Hi, Ed! Root path"}`
// 	json.NewEncoder(w).Encode(contents)
// }

// GetEvents Get the event
// func GetEvents(w http.ResponseWriter, r *http.Request) {

// 	// a map of values
// 	params := mux.Vars(r)

// 	fmt.Printf("Params: %v\n", params)

// 	fmt.Printf("val is: '%s'\n", params["val"])

// 	queryResults := query.GetBodyQueryByEventsResty("darkstar", 9200, "mybrooklyn", params["val"], 0, 50)

// 	objectResults, _ := query.ParseElasticJSONObject(queryResults)

// 	json.NewEncoder(w).Encode(objectResults.Hits.Hits)

// }

// GetWeather Get the event
// func GetWeather(w http.ResponseWriter, r *http.Request) {

// 	// a map of values
// 	params := mux.Vars(r)

// 	fmt.Printf("Params: %v\n", params)

// 	fmt.Printf("val is: '%s'\n", params["val"])

// 	queryResults := query.GetByWeather("darkstar", 9200, "mybrooklyn", params["val"], 0, 50)

// 	objectResults, _ := query.ParseElasticJSONObject(queryResults)

// 	json.NewEncoder(w).Encode(objectResults.Hits.Hits)

// }

// GetDbWeather Get the event
func GetDbWeather(w http.ResponseWriter, r *http.Request) {

	// a map of values
	params := mux.Vars(r)

	// []SimpleBrooklyn

	rows := DBDoIt(params["val"])

	json.NewEncoder(w).Encode(rows)

}

// Server does what
// func Server() {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/", GetRootPath).Methods("GET")
// 	router.HandleFunc("/getEvents/{val}", GetEvents).Methods("GET")
// 	router.HandleFunc("/getWeather/{val}", GetWeather).Methods("GET")
// 	router.HandleFunc("/getDbWeather/{val}", GetDbWeather).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":8080", router))

// }

// DBDoIt access via a database
func DBDoIt(findEvent string) []SimpleBrooklyn {

	// spring.datasource.url: jdbc:mysql://localhost:3306/dblearning
	// spring.datasource.username: edwinbrown
	// spring.datasource.password: userpassword

	db, err := sql.Open("mysql",
		"edwinbrown:userpassword@tcp(127.0.0.1:3306)/dblearning")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeing if database is available")
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	var (
		startDate     string
		weather       string
		temperature   float32
		precipitation float32
		events        sql.NullString
	)

	// rows, err := db.Query("select start_date, weather,temperature,precipitation,events from brooklyn_bridge")

	rows, err := db.Query("select start_date, weather,temperature,precipitation,events from brooklyn_bridge where events = ?", findEvent)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	sbrows := make([]SimpleBrooklyn, 1)

	var sb SimpleBrooklyn

	for rows.Next() {
		err := rows.Scan(&startDate, &weather, &temperature, &precipitation, &events)
		if err != nil {
			log.Fatal(err)
		}

		if events.Valid {

			sb.Events = events.String
			sb.Precipitation = precipitation
			sb.StartDate = startDate
			sb.Temperature = temperature
			sb.Weather = weather

			log.Printf("date: %s, weather: %s, temp: %f, precip: %f, events: %s\n", startDate, weather, temperature, precipitation, events.String)
		} else {

			sb.Events = ""
			sb.Precipitation = precipitation
			sb.StartDate = startDate
			sb.Temperature = temperature
			sb.Weather = weather

			log.Printf("date: %s, weather: %s, temp: %f, precip: %f, events: \n", startDate, weather, temperature, precipitation)

		}

		sbrows = append(sbrows, sb)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Done")

	return sbrows
}

// databaseCalls the database call
func databaseCalls() {
	DBDoIt("Halloween")
}

// serverInfo server information
func serverInfo(uri string) {
	// get the information about the cluster
	status, _ := query.GetElasticServerInfo(uri)

	fmt.Printf("Name: %s\nVersion: %s\nLucene Version: %s\nCluster name: %s\nTagline: %s\n", status.Name, status.Version.Number,
		status.Version.LuceneVersion, status.ClusterName, status.Tagline)

}

// elasticCalls the elastic calls
func elasticCalls() {

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

}

func main() {
	elasticURI := flag.String("elasticUri", "", "The URI of ElasticSearch")
	dbFlag := flag.Bool("database", false, "Use database for calls")
	elastic := flag.Bool("elastic", false, "Use elastic for calls")
	server := flag.Bool("restServer", false, "Use RESTful server")

	flag.Parse()

	fmt.Printf("elasticURI: '%s'\n", *elasticURI)

	if *elasticURI != "" {
		serverInfo(*elasticURI)
	}

	if *dbFlag {
		databaseCalls()
	}

	if *elastic {
		elasticCalls()
	}

	if *server {
		webbie.Server()
		// Server()

	}
	// }
}
