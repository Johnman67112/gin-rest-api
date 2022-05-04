package database

import (
	"log"

	"github.com/Johnman67112/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnect() {
	conn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conn))
	if err != nil {
		log.Panic("Database connection error")
	}

	DB.AutoMigrate(&models.Student{})
}
