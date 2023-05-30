package order_repository

import (
	"assignment-2/entity"
	"assignment-2/pkg/errs"
)

type OrderRepository interface {
	CreateOrder(orderPayload *entity.Order, itemPayload []entity.Item) (*entity.Order, errs.MessageErr)
	GetAllData() ([]entity.Order, errs.MessageErr)
	UpdateOrderById(orderID uint, orderPayload *entity.Order, itemPayload []entity.Item) (*entity.Order, errs.MessageErr)
	DeleteOrderById(orderID uint) errs.MessageErr
}
