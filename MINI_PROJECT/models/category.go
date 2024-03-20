package models

import "gorm.io/gorm"

type Categories struct {
	ID          int        `gorm:"primary_key;autoIncrement" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Products    []Products `gorm:"foreignKey:CategoryID" json:"products"`
}

func MigrateCategories(db *gorm.DB) error {
	err := db.AutoMigrate(&Categories{})
	return err

}
