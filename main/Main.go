package main

import (
	"fmt"
	"os"

	"github.com/EdSwArchitect/brooklyn-es/query"
)

func main() {
	args := os.Args

	fmt.Printf("Hi, Ed: %v\n", args)

	uri := args[1]

	fmt.Printf("The one arg is: '%s'\n", uri)

	// if esp, err := query.ReadJSONObject("/home/edbrown/Documents/brooklyn_results.json"); err == nil {

	// 	fmt.Printf("Took - %d: TimedOut - %t \n", esp.Took, esp.TimedOut)
	// 	fmt.Printf("Shards.total: %d Shards.successful: %d Shards.skipped: %d Shards.failed :%d\n",
	// 		esp.Shards.Total, esp.Shards.Successful, esp.Shards.Skipped, esp.Shards.Failed)

	// 	fmt.Printf("Shards.hits.total: %d Shards.hits.maxScore: %f\n", esp.Hits.Total, esp.Hits.MaxScore)
	// 	fmt.Printf("hits.total: %d\n", esp.Hits.Total)
	// 	fmt.Printf("hits.total length: %d\n", len(esp.Hits.Hits))

	// 	for _, dahit := range esp.Hits.Hits {
	// 		fmt.Printf("\t*** time: %s location: %s lat/lon: %f/%f weather:%s \n", dahit.Source.Timestamp,
	// 			dahit.Source.Location, dahit.Source.Lat,
	// 			dahit.Source.Lon, dahit.Source.WeatherSummary)
	// 	}
	// } else {
	// 	fmt.Println("Parse failed")
	// }

	status, _ := query.GetElasticServerInfo(uri)

	fmt.Printf("Name: %s\nVersion: %s\nLucene Version: %s\nCluster name: %s\nTagline: %s\n", status.Name, status.Version.Number,
		status.Version.LuceneVersion, status.ClusterName, status.Tagline)

	// queryResuts := query.GetQueryResults("darkstar", 9200, "mybrooklyn")

	// fmt.Printf("Query results\n%s\n\n*********\n\n", queryResuts)

	// queryResuts = query.GetLimitedQueryResults("darkstar", 9200, "mybrooklyn", 7100, 50)

	// fmt.Printf("Query results\n%s\n", queryResuts)

	queryResuts := query.GetBodyQueryByEventsResty("darkstar", 9200, "mybrooklyn", "Father's", 0, 50)

	fmt.Printf("\n\nFLAG FLAG FLAG\n\nQuery results\n%s\n", queryResuts)

	//Independence Day

}
