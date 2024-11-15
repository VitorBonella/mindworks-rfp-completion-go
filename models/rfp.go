package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type RFP struct{
	Id uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"-"`
	Name string `json:"name" gorm:"unique" validate:"required"`
	CreationDate *time.Time `json:"creation_date"`
	EndDate *time.Time `json:"end_date"`
	Status string `json:"status"`
	Requirements []Requirement `json:"requirements" gorm:"foreignKey:RFPId;references:Id"`
	Equipments []Equipment `json:"equipments" gorm:"foreignKey:id;many2many:rfp_equipment;"`
}

type Requirement struct{
	Id uint `json:"id" gorm:"primaryKey"`
	Requirement string `json:"requirement"`
	RFPId uint `json:"rfp_id"`
}

const RFPStatusCreated = "Created"
const RFPtatusProcessing = "Processing"
const RFPStatusFinished = "Finished"
const RFPStatusFinishedWithError = "Finished With Error"

func NewRFP(name string, requirements []string, equipments []Equipment, userId uint) (*RFP,error){

	if name == ""{
		return nil, errors.New("invalid name")
	}

	if len(requirements) == 0{
		return nil, errors.New("empty requirements")
	}

	creationDate := time.Now()

	var rfpList []Requirement
	for _ , v := range requirements{
		rfpList = append(rfpList, Requirement{Requirement: v})
	}

	rfp := RFP{
		Name: name,
		Requirements: rfpList,
		Equipments: equipments,
		UserId: userId,
		CreationDate: &creationDate,
		Status: RFPStatusCreated,
	}

	return &rfp, nil
}

type Question struct {
	Question    string `json:"question"`
	Answer      string `json:"answer"`
	Source      string `json:"source"`
	Description string `json:"description"`
}

func GenerateQuestionJSON(requirements []Requirement) ([]string, error) {
	var result []string
	blockSize := 50
	total := len(requirements)

	for i := 0; i < total; i += blockSize {
		end := i + blockSize
		if end > total {
			end = total
		}

		block := make(map[string]Question)
		for j, req := range requirements[i:end] {
			block[fmt.Sprintf("QUESTION_%d", i+j+1)] = Question{
				Question:    req.Requirement,
				Answer:      "", // Default or dynamic value
				Source:      "", // Default or dynamic value
				Description: "", // Default or dynamic value
			}
		}

		jsonData, err := json.Marshal(block)
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonData))
	}

	return result, nil
}