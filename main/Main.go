package main

import (
	"fmt"

	"github.com/EdSwArchitect/brooklyn-es/query"
)

func main() {
	fmt.Println("Hi, Ed")

	if esp, err := query.ReadJSONObject("/home/edbrown/Documents/brooklyn_results.json"); err == nil {

		fmt.Printf("Took - %d: TimedOut - %t \n", esp.Took, esp.TimedOut)
		fmt.Printf("Shards.total: %d Shards.successful: %d Shards.skipped: %d Shards.failed :%d\n",
			esp.Shards.Total, esp.Shards.Successful, esp.Shards.Skipped, esp.Shards.Failed)

		fmt.Printf("Shards.hits.total: %d Shards.hits.maxScore: %f\n", esp.Hits.Total, esp.Hits.MaxScore)
		fmt.Printf("hits.total: %d\n", esp.Hits.Total)
		fmt.Printf("hits.total length: %d\n", len(esp.Hits.Hits))

		for _, dahit := range esp.Hits.Hits {
			fmt.Printf("\t*** time: %s location: %s lat/lon: %f/%f weather:%s \n", dahit.Source.Timestamp,
				dahit.Source.Location, dahit.Source.Lat,
				dahit.Source.Lon, dahit.Source.WeatherSummary)
		}
	} else {
		fmt.Println("Parse failed")
	}

}
