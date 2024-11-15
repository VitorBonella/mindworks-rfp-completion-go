package database

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
	"fmt"
	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
  )

func NewDBConnection() (*gorm.DB, error){
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil{
		fmt.Println("Error connecting to database")
		return nil, err
	}

	db.AutoMigrate(&models.User{},&models.Equipment{},&models.RFP{},&models.Requirement{})

	return db, nil
}