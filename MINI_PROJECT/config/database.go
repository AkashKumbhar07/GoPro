package config

import (
	"log"
	"os"

	"github.com/YohanJirage/Go_Lang/models"
	"github.com/YohanJirage/Go_Lang/storage"
	"gorm.io/gorm"
)

func DB_Connect() (*gorm.DB, error) {

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBname:   os.Getenv("DB_NAME"),
		SSLModes: os.Getenv("DB_SSLMODE "),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("FAILED TO CREATE NEW CONNECTION")
	}
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.Categories{})
	db.AutoMigrate(&models.Products{})
	db.AutoMigrate(&models.Payments{})
	db.AutoMigrate(&models.Orders{})
	db.AutoMigrate(&models.Carts{})

	return db, nil

}
