package models

import "gorm.io/gorm"


type Users struct {
    ID       int       `gorm:"primary_key;autoIncrement" json:"id"`
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    Addresses []Address `gorm:"foreignKey:UserID" json:"addresses"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err

}

