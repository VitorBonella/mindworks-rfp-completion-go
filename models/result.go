package models

import "time"

type Result struct{
	Id uint `json:"id" gorm:"primaryKey"`
	RFPId uint `json:"rfp_id"`
	EquipmentId uint `json:"equipment_id"`
	Text string `json:"text"`
	EndDate *time.Time `json:"end_date"`
}