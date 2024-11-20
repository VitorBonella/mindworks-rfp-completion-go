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
	defer CloseDBConnection(db)

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
	defer CloseDBConnection(db)

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
	defer CloseDBConnection(db)

	var rfps []*models.RFP

	err = db.Preload("Equipments").Preload("Requirements").Where("user_id = ?",userId).Order("creation_date DESC").Find(&rfps).Error
	if err != nil{
		log.Println(err)
		return nil,err
	}

	return rfps,nil
}

func GetNewestCreatedRFP() (*models.RFP, error){

	db, err := NewDBConnection()
	if err != nil{
		return nil,err
	}
	defer CloseDBConnection(db)

	var newestRFP *models.RFP

	err = db.Preload("Equipments").Preload("Requirements").Where("status = ?",models.RFPStatusCreated).Order("creation_date DESC").First(&newestRFP).Error
	if err != nil{
		return nil,err
	}

	return newestRFP,nil

}


func ListProcessingRFP() ([]*models.RFP, error){

	db, err := NewDBConnection()
	if err != nil{
		return nil,err
	}
	defer CloseDBConnection(db)

	var processingRFPs []*models.RFP

	err = db.Preload("Equipments").Preload("Requirements").Where("status = ?",models.RFPtatusProcessing).Find(&processingRFPs).Error
	if err != nil{
		return nil,err
	}

	return processingRFPs,nil

}

func SetRFPStatus(rfp *models.RFP, status string) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}
	defer CloseDBConnection(db)

	rfp.Status = status
    
	err = db.Updates(rfp).Error
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}