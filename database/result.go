package database

import (
	"log"

	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
)


func CreateResult(r *models.Result) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}
	defer CloseDBConnection(db)

	err = db.Create(&r).Error
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}

func GetResults(rfpId uint, equipId uint) ([]*models.Result,error){

	db, err := NewDBConnection()
	if err != nil{
		return nil,err
	}
	defer CloseDBConnection(db)

	var results []*models.Result

	err = db.Where("rfp_id = ? and equipment_id = ?",rfpId,equipId).Find(&results).Error
	if err != nil{
		log.Println(err)
		return nil,err
	}

	return results, nil
}