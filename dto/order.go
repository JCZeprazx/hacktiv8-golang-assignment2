package dto

import (
	"assignment-2/entity"
	"time"
)

type OrderRequest struct {
	OrderedAt    time.Time
	CustomerName string
	Items        []ItemRequest
}

type OrderResponse struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CustomerName string
	Items        []ItemResponse
}

type NewOrderResponse struct {
	StatusCode int
	Message    string
	Data       OrderRequest
}

type GetAllOrderResponse struct {
	StatusCode int
	Message    string
	Data       []OrderResponse
}

type UpdateOrderById struct {
	StatusCode int
	Message    string
	Data       OrderResponse
}

type DeleteOrderById struct {
	StatusCode int
	Message    string
}

func (o *OrderRequest) OrderRequestToEntity() *entity.Order {
	return &entity.Order{
		CustomerName: o.CustomerName,
		OrderedAt:    o.OrderedAt,
	}
}
