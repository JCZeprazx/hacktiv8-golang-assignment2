package entity

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `gorm:"not null; unique"`
	Description string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	OrderID     int
}
