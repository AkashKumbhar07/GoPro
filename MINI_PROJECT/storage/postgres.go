package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBname   string
	SSLModes string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Password, config.User, config.DBname, config.SSLModes)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn, // Your data source name (DSN)
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
