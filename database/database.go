package database

import (
	"crudProject/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	Database = DbInstance{Db: db}
}
