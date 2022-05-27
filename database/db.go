package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Johnman67112/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnect() {
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DBPORT")
	ssl := os.Getenv("SSL")

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, dbport, ssl)
	DB, err = gorm.Open(postgres.Open(conn))
	if err != nil {
		log.Panic("Database connection error")
	}

	DB.AutoMigrate(&models.Student{})
}
