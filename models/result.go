package models

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type Result struct{
	Id uint `json:"id" gorm:"primaryKey"`
	RFPId uint `json:"rfp_id"`
	EquipmentId uint `json:"equipment_id"`
	Text string `json:"text"`
	EndDate *time.Time `json:"end_date"`
}


type QuestionMap struct{
	Map map[string]Question
}

func ConcatResults(results []*Result) *QuestionMap {
	// Initialize a new map to hold the combined results
	combinedMap := make(map[string]Question)

	for _, r := range results {
		// Temporary map to hold the unmarshalled data
		text := strings.Replace(r.Text,"```","",-1)
		qmap := make(map[string]Question)
		err := json.Unmarshal([]byte(text), &qmap)
		if err != nil {
			log.Println("Error unmarshalling JSON:", err)
			return nil
		}

		// Merge qmap into combinedMap
		for key, question := range qmap {
			// Check for key conflicts if needed (optional logic)
			// combinedMap[key] will overwrite by default, but you can customize
			combinedMap[key] = question
		}
	}

	return &QuestionMap{Map: combinedMap}
}