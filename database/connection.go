package database

import (
	"fmt"

	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() (*gorm.DB, error){
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil{
		fmt.Println("Error connecting to database",err)
		return nil, err
	}

	db.AutoMigrate(&models.User{},&models.Equipment{},&models.RFP{},&models.Requirement{},&models.Result{})

	return db, nil
}