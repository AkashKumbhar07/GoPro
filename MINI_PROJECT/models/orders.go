package models

import "gorm.io/gorm"

type Orders struct {
	OrderID   int `gorm:"primary_key;autoIncrement" json:"order_id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	PaymentID int `json:"payment_id"`
	AddressID int `json:"address_id"`
}

func MigrateOrders(db *gorm.DB) error {
	err := db.AutoMigrate(&Orders{})
	return err

}
