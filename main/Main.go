package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/EdSwArchitect/brooklyn-es/query"
)

// Source for testing
type Source struct {
	TowardsBrooklyn  int32   `json:"towards_brooklyn"`
	Location         string  `json:"location"`
	Lat              float64 `json:"lat"`
	TowardsManhattan int32   `json:"towards_manhattan"`
	Timestamp        string  `json:"@timestamp"`
	WeatherSummary   string  `json:"weather_summary"`
	Temperature      int32   `json:"temperature"`
	Events           string  `json:"events"`
	Pedestriants     int32   `json:"Pedestrians"`
	Lon              float64 `json:"lon"`
	Precipitation    float64 `json:"precipitation"`
}

// Shards stuff
type Shards struct {
	Total      int32 `json:"total"`
	Successful int32 `json:"successful"`
	Skipped    int32 `json:"skipped"`
	Failed     int32 `json:"failed"`
}

// BigHits hits
type BigHits struct {
	Total    int32   `json:"total"`
	MaxScore float32 `json:"max_score"`
	Hits     []Hits  `json:"hits"`
}

// Hits type
type Hits struct {
	Index   string  `json:"_index"`
	DocType string  `json:"_type"`
	Id      string  `json:"_id"`
	Score   float32 `json:"_score"`
	Source  Source  `json:"_source"`
}

// ESInfo stuff
type ESInfo struct {
	Took     int32   `json:"took"`
	TimedOut bool    `json:"timed_out"`
	Shards   Shards  `json:"_shards"`
	Hits     BigHits `json:"hits"`
}

func main() {
	fmt.Println("Hi, Ed")

	contents := query.ReadJSON("/home/edbrown/Documents/brooklyn_results.json")

	fmt.Println(contents)

	dec := json.NewDecoder(strings.NewReader(contents))

	for {
		var esp ESInfo

		if err := dec.Decode(&esp); err == io.EOF {
			fmt.Println("EOF")
			break
		} else if err != nil {
			log.Fatal(err)
		}

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
	}

}
