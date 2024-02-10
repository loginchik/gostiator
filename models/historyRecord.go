package models

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"log"
	"os"
	"sort"
	"time"
)

type HistoryRecords struct {
	Records []*HistoryRecord
}

type HistoryRecord struct {
	ID         int
	DateStamp  time.Time `json:"date_stamp"`
	DateString string    `json:"date_string"`
	RecordType string    `json:"record_type"`
	Content    string    `json:"content"`
}

var historyFile = "history.json"

func (hr *HistoryRecord) Save() {
	var records = GetHistoryRecords()

	if hr.DateString == "" {
		hr.DateString = time.Time.Format(hr.DateStamp, "02.01.2006")
	}
	if hr.ID == 0 {
		hr.ID = len(records.Records) + 1
	}

	records.Records = append(records.Records, hr)
	fileData, _ := json.Marshal(&records)
	saveError := os.WriteFile(historyFile, fileData, 0666)
	if saveError != nil {
		log.Println("Error saving json", saveError)
	} else {
		log.Println("Saved new history record")
	}
}

func GetHistoryRecords() *HistoryRecords {
	var records = new(HistoryRecords)

	file, err := fyne.LoadResourceFromPath(historyFile)
	if err != nil {
		_, err = os.Create(historyFile)
		if err != nil {
			panic("")
		}
		return records
	}
	jsonErr := json.Unmarshal(file.Content(), &records)
	if jsonErr != nil {
		log.Println("Error unmarshalling json from get all", jsonErr)
	}

	if len(records.Records) > 0 {
		sort.Slice(records.Records, func(i, j int) bool {
			firstDateString := records.Records[i].DateStamp
			lastDateString := records.Records[j].DateStamp
			return firstDateString.After(lastDateString)
		})
	}

	log.Printf("Read %d history records from json", len(records.Records))
	return records
}
