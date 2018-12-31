package query

import (
	"fmt"
	"log"
	"strings"

	"encoding/json"
	"net/url"
	"strconv"

	resty "gopkg.in/resty.v1"
)

// ElasticServerInfo JSON structure
type ElasticServerInfo struct {
	Name        string  `json:"name"`
	ClusterName string  `json:"cluster_name"`
	ClusterUUID string  `json:"cluster_uuid"`
	Version     Version `json:"version"`
	Tagline     string  `json:"tagline"`
}

// Version JSON object
type Version struct {
	Number                           string `json:"number"`
	BuildFlavor                      string `json:"build_flavor"`
	BuildType                        string `json:"build_type"`
	BuildHash                        string `json:"build_hash"`
	BuildDate                        string `json:"build_date"`
	BuildSnapshot                    bool   `json:"build_snapshot"`
	LuceneVersion                    string `json:"lucene_version"`
	MinimumWireCompatibilityVersion  string `json:"minimum_wire_compatibility_version"`
	MinimumIndexCompatibilityVersion string `json:"minimum_index_compatibility_version"`
}

// GetElasticServerInfo get the ElasticSearch status
func GetElasticServerInfo(url string) (status ElasticServerInfo, err error) {
	resp, err := resty.R().Get(url)

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())

	dec := json.NewDecoder(strings.NewReader(resp.String()))

	if err := dec.Decode(&status); err != nil {
		log.Fatal(err)
		return status, err
	}

	fmt.Printf("\nElastic object: %v\n", status)

	return status, err
}

// GetQueryResults get the results of an ElasticSearch query
func GetQueryResults(host string, port int, index string) string {
	url := "http://" + host + ":" + strconv.Itoa(port) + "/" + index + "/_search?pretty=true"

	fmt.Printf("The Url: '%s'\n", url)

	resp, _ := resty.R().Get(url)

	return resp.String()
}

// GetLimitedQueryResults get the results of an ElasticSearch query
func GetLimitedQueryResults(host string, port int, index string, start int, size int) string {
	url := "http://" + host + ":" + strconv.Itoa(port) + "/" + index + "/_search?pretty=true" + "&from=" + strconv.Itoa(start) + "&size=" + strconv.Itoa(size)

	fmt.Printf("The Url: '%s'\n", url)

	resp, _ := resty.R().Get(url)

	return resp.String()
}

// GetLimitedQueryByEvents get the results of an ElasticSearch query
func GetLimitedQueryByEvents(host string, port int, index string, event string, start int, size int) string {
	uri := "http://" + host + ":" + strconv.Itoa(port) + "/" + index + "/_search?pretty=true"
	uri += "&from=" + strconv.Itoa(start) + "&size=" + strconv.Itoa(size) + "&q=events:"

	uri += url.QueryEscape(event)

	fmt.Printf("The Url: '%s'\n", uri)

	resp, _ := resty.R().Get(uri)

	return resp.String()
}

// GetLimitedQueryByEventsResty get the results of an ElasticSearch query
func GetLimitedQueryByEventsResty(host string, port int, index string, event string, start int, size int) string {
	// uri := "http://" + host + ":" + strconv.Itoa(port) + "/" + index + "/_search?pretty=true"
	// uri += "&from=" + strconv.Itoa(start) + "&size=" + strconv.Itoa(size) + "&q=events:"

	// uri += url.QueryEscape(event)

	// fmt.Printf("The Url: '%s'\n", uri)

	// resp, _ := resty.R().Get(uri)

	portion := url.QueryEscape(event)

	resp2, err := resty.R().
		SetQueryParams(map[string]string{
			"pretty":                       "true",
			"allow_partial_search_results": "true",
			"from":                         strconv.Itoa(start),
			"size":                         strconv.Itoa(size),
			"q":                            "events:" + portion,
		}).Get("http://" + host + ":" + strconv.Itoa(port) + "/" + index + "/_search?")

	if err != nil {
		fmt.Printf("Some error: %v\n", err)
	}

	return resp2.String()
}

// GetBodyQueryByEventsResty get the results of an ElasticSearch query
func GetBodyQueryByEventsResty(host string, port int, index string, event string, start int, size int) string {

	body := `{
		"query" : {
			"match" :  { 
				"events" : "` + event + `" 
			}
		},
		"from" : ` + strconv.Itoa(start) + `,
		"size" : ` + strconv.Itoa(size) + `

	}
	`

	fmt.Printf("The body\n%s\n", body)

	resp2, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetQueryParams(map[string]string{
			"pretty": "true"}).
		Post("http://" + host + ":" + strconv.Itoa(port) + "/" + index + "/_search")

	if err != nil {
		fmt.Printf("Some error: %v\n", err)
	}

	return resp2.String()
}

// GetByWeather get the results of an ElasticSearch query
func GetByWeather(host string, port int, index string, weather string, start int, size int) string {

	body := `{
		"query" : {
			"match" :  { 
				"weather_summary" : "` + weather + `" 
			}
		},
		"from" : ` + strconv.Itoa(start) + `,
		"size" : ` + strconv.Itoa(size) + `

	}
	`

	fmt.Printf("The body\n%s\n", body)

	resp2, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetQueryParams(map[string]string{
			"pretty": "true"}).
		Post("http://" + host + ":" + strconv.Itoa(port) + "/" + index + "/_search")

	if err != nil {
		fmt.Printf("Some error: %v\n", err)
	}

	return resp2.String()
}
