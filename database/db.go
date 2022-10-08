package database

import (
	"api-rest-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDatabase() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados", err.Error())
	}

	DB.AutoMigrate(&models.Student{})
}
