package database

import "github.com/VitorBonella/mindworks-rfp-completion-go/models"

func CreateEquipment(equipment *models.Equipment) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}

	db.Create(&equipment)

	return nil
}

func UpdateEquipment(equipment *models.Equipment) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}

	db.Updates(&equipment)

	return nil
}

func DeleteEquipment(equipment *models.Equipment) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}

	db.Debug().Delete(&equipment)

	return nil
}

func ListEquipment(userId uint) ([]*models.Equipment,error) {

	db, err := NewDBConnection()
	if err != nil{
		return nil, err
	}

	var equipments []*models.Equipment

	db.Where("user_id = ?",userId).Find(&equipments)

	return equipments,nil
}