package config

import (
	"log"
	"os"

	"github.com/YohanJirage/Go_Lang/models"
	"github.com/YohanJirage/Go_Lang/storage"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DB_Connect() (*gorm.DB, error) {

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBname:   os.Getenv("DB_NAME"),
		SSLModes: os.Getenv("DB_SSLMODE "),
	}

	Db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("FAILED TO CREATE NEW CONNECTION")
	}
	Db.AutoMigrate(&models.Users{})
	Db.AutoMigrate(&models.Address{})
	Db.AutoMigrate(&models.Categories{})
	Db.AutoMigrate(&models.Products{})
	Db.AutoMigrate(&models.Payments{})
	Db.AutoMigrate(&models.Orders{})
	Db.AutoMigrate(&models.Carts{})

	return Db, nil

}
