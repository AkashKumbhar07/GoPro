package models

import (
	"gorm.io/gorm"
)

type Products struct {
	ID          int     `gorm:"primary_key;autoIncrement" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"-"`
}

func MigrateProducts(db *gorm.DB) error {
	err := db.AutoMigrate(&Products{})
	return err

}
