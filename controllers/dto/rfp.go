package dto

import "github.com/VitorBonella/mindworks-rfp-completion-go/models"

type RFP struct{

	Name string
	Requirements []string
	Equipments []models.Equipment
}