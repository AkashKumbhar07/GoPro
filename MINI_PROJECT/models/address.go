package models

import "gorm.io/gorm"

type Address struct {
	ID       int    `gorm:"primary_key;autoIncrement" json:"id"`
	Text     string `json:"text"`
	District string `json:"district"`
	State    string `json:"state"`
	Pincode  string `json:"pincode"`
	UserID   int    `json:"-"`
}

func MigrateAddress(db *gorm.DB) error {
	err := db.AutoMigrate(&Address{})
	return err

}
