package models

import "gorm.io/gorm"

type Carts struct {
	CartID    int      `gorm:"primary_key;autoIncrement" json:"cart_id"`
	UserID    int      `json:"user_id"`
	ProductID int      `json:"product_id"`
	Quantity  int      `json:"quantity"`
	Product   Products `gorm:"foreignKey:ProductID"`
}

func MigrateCarts(db *gorm.DB) error {
	err := db.AutoMigrate(&Carts{})
	return err

}