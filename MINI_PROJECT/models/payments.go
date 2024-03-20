package models

import "gorm.io/gorm"

type Payments struct {
	PaymentID int      `gorm:"primary_key;autoIncrement" json:"payment_id"`
	Amount    float64  `json:"amount"`
	ProductID int      `json:"product_id"`
	UserID    int      `json:"user_id"`
	Product   Products `gorm:"foreignKey:ProductID"`
}

func MigratePayments(db *gorm.DB) error {
	err := db.AutoMigrate(&Payments{})
	return err

}
