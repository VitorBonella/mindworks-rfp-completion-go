package database

import (
	"log"

	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
)

func CreateRFP(rfp *models.RFP) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}

	for i := range rfp.Equipments{
		rfp.Equipments[i].UserId = rfp.UserId
	}

	err = db.Create(&rfp).Error
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}

func GetRFP(rfpId uint) (*models.RFP,error){

	db, err := NewDBConnection()
	if err != nil{
		return nil,err
	}

	var rfp *models.RFP

	err = db.Preload("Equipments").Preload("Requirements").Where("id = ?",rfpId).First(&rfp).Error
	if err != nil{
		log.Println(err)
		return nil,err
	}

	return rfp,nil

}

func ListRFP(userId uint) ([]*models.RFP,error){

	db, err := NewDBConnection()
	if err != nil{
		return nil,err
	}

	var rfps []*models.RFP

	err = db.Preload("Equipments").Preload("Requirements").Where("user_id = ?",userId).Find(&rfps).Error
	if err != nil{
		log.Println(err)
		return nil,err
	}

	return rfps,nil
}

func ListNewestCreatedRFP() (*models.RFP, error){

	db, err := NewDBConnection()
	if err != nil{
		return nil,err
	}

	var newestRFP *models.RFP

	err = db.Preload("Equipments").Preload("Requirements").Where("status = ?",models.RFPStatusCreated).Order("creation_date DESC").First(&newestRFP).Error
	if err != nil{
		return nil,err
	}

	return newestRFP,nil

}

func SetRFPStatus(rfp *models.RFP, status string) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}

	rfp.Status = status
    
	err = db.Updates(rfp).Error
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}