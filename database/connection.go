package database

import (
	"fmt"

	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() (*gorm.DB, error){
	db, err := gorm.Open(sqlite.Open("./data/gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil{
		fmt.Println("Error connecting to database",err)
		return nil, err
	}
	
	db.AutoMigrate(&models.User{},&models.Equipment{},&models.RFP{},&models.Requirement{},&models.Result{})

	return db, nil
}

func CloseDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB() // Retrieve the underlying *sql.DB
	if err != nil {
		return fmt.Errorf("failed to retrieve database instance: %w", err)
	}

	return sqlDB.Close() // Close the connection
}