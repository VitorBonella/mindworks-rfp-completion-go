package database

import "github.com/VitorBonella/mindworks-rfp-completion-go/models"

func CreateEquipment(equipment *models.Equipment) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}
	defer CloseDBConnection(db)

	err = db.Create(&equipment).Error
	if err != nil{
		return err
	}

	return nil
}

func UpdateEquipment(equipment *models.Equipment) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}
	defer CloseDBConnection(db)

	db.Updates(&equipment)

	return nil
}

func DeleteEquipment(equipment *models.Equipment) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}
	defer CloseDBConnection(db)

	db.Debug().Delete(&equipment)

	return nil
}

func ListEquipment(userId uint) ([]*models.Equipment,error) {

	db, err := NewDBConnection()
	if err != nil{
		return nil, err
	}
	defer CloseDBConnection(db)

	var equipments []*models.Equipment

	db.Where("user_id = ?",userId).Find(&equipments)

	return equipments,nil
}