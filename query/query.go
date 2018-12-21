package query

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

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
