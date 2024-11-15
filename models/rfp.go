package models

import (
	"errors"
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

const RFTStatusCreated = "Created"

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
		Status: RFTStatusCreated,
	}

	return &rfp, nil
}