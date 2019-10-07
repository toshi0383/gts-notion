package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RaceRecord struct {
	CourseCode string `json:"course_code"`
	Tire       int    `json:"tire"`
	Fuel       int    `json:"fuel"`
	Category   string `json:"category"`
	Fav        string `json:"fav"` // Yes or No
	Laps       int    `json:"laps"`
	Date       string `json:"date"`
	Tires      []int  `json:"tires"`
	Series     int    `json:"series"`
}

func (r *RaceRecord) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PropertyID string

func SubmitTransaction(record RaceRecord) {

	url := "https://www.notion.so/api/v3/submitTransaction"

	var jsonStr = []byte(`{"operations":[{"id":"618cc6a9-9fc8-49a5-93c3-d19b2ba87b86","table":"block","path":["properties","AAXJ"],"command":"set","args":[["Yes"]]}]}`)
	body := bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("cookie", "__cfduid=d5c6ddf24bf4a816d7e8789c5e076cfae1570105307")
	req.Header.Set("cookie", "token_v2=e7dfc7e538ff82c639ec2ff74725905b6b2e16e6057fd5a1c855a3dd18e226df12b8f301b321a1aea7ef436cc35dc65e6ad7569891054604a8bb9d1e17c0adc53bcc2e241a30cdc88c3969d1a486")

	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
}
