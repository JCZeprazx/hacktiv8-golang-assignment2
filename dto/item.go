package dto

import (
	"assignment-2/entity"
	"time"
)

type ItemRequest struct {
	ItemCode    string
	Description string
	Quantity    int
}

type ItemResponse struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ItemCode    string
	Description string
	Quantity    int
	OrderID     int
}

func (i *ItemRequest) ItemRequestToEntity() *entity.Item {
	return &entity.Item{
		ItemCode:    i.ItemCode,
		Description: i.Description,
		Quantity:    i.Quantity,
	}
}
