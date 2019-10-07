package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/toshi0383/gts-notion/gtsport"
	"github.com/toshi0383/gts-notion/notion"
)

func main() {
	events := gtsport.GetManufacturerEventRounds()
	var records []notion.RaceRecord

	// TODO: run in parallel to improve performance
	for _, e := range events {
		race := e.Race
		reg := e.Regulation

		t := e.BeginDate
		beginDate := fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())

		record := notion.RaceRecord{
			CourseCode: e.Track.CourseCode,
			Tire:       race.ConsumeTire,
			Fuel:       race.ConsumeFuel,
			Category:   reg.CarCategoryTypes[0],
			Fav:        "No",
			Laps:       race.RaceLimitLaps,
			Date:       beginDate,
			Tires:      []int{reg.LimitTireF, reg.NeedTireF},
			Series:     0, // manu
		}
		records = append(records, record)
	}

	bytes, err := json.Marshal(records)
	if err != nil {
		log.Fatalf("Marshal failed: %v\n", err)
	}
	os.Stdout.Write(bytes)
}
