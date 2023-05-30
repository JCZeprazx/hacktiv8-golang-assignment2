package entity

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	OrderedAt    time.Time `gorm:"default:now()"`
	CustomerName string    `gorm:"not null"`
	Items        []Item
}
