package database

import "github.com/VitorBonella/mindworks-rfp-completion-go/models"

func CreateUser(user *models.User) error{

	db, err := NewDBConnection()
	if err != nil{
		return err
	}

	db.Create(&user)

	return nil
}

func GetUserByName(name string) (*models.User,error){

	db, err := NewDBConnection()
	if err != nil{
		return nil,err
	}

	var user *models.User

	db.Where("name = ?", name).First(&user)

	return user,nil
}

func GetUserById(id string) (*models.User,error){

	db, err := NewDBConnection()
	if err != nil{
		return nil,err
	}

	var user *models.User

	db.Where("id = ?", id).First(&user)

	return user,nil
}