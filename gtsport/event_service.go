package gtsport

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func requestWithBody(body io.Reader) (*http.Request, error) {
	url := "https://www.gran-turismo.com/jp/api/gt7sp/event/"
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	return req, nil
}

type Series int

const (
	manu Series = iota
	nations
)

func GetManufacturerEventRounds() []GameParameterEvent {
	return getRounds(manu)
}

func getRounds(series Series) []GameParameterEvent {

	var jsonStr = []byte("job=10")
	req, err := requestWithBody(bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	si, err := UnmarshalSeasonInfo(body)
	seasons := si.SeasonInformation.Seasons

	var index int
	switch series {
	case manu:
		index = len(seasons) - 1
	case nations:
		index = len(seasons) - 2
	}

	rounds := seasons[index].Rounds

	var events []GameParameterEvent

	for _, r := range rounds {
		bodystr := fmt.Sprintf("event_id_csv=%s&job=1", r.EventID)
		var jsonStr = []byte(bodystr)
		req, err := requestWithBody(bytes.NewBuffer(jsonStr))
		if err != nil {
			panic(err)
		}
		var PTransport http.RoundTripper = &http.Transport{Proxy: http.ProxyFromEnvironment}
		client := http.Client{Transport: PTransport}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		ei, err := UnmarshalEventInfo(body)
		if err != nil {
			panic(err)
		}
		if len(ei.Event) == 0 {
			continue
		}
		fe := ei.Event[0]
		ve := fe.Value[0]
		event := ve.ValueClass.GameParameter.Events[0]
		track := ve.ValueClass.GameParameter.Tracks[0]
		event.BeginDate = r.StartDate
		event.Track = track
		events = append(events, event)
	}

	return events
}
