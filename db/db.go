package db

import (
	"log"

	"github.com/matheusrf96/gin-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=ginrest port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Cannot connect to databse")
	}

	db.AutoMigrate(&models.Student{})
}
