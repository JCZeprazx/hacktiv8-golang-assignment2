package order_postgres

import (
	"assignment-2/entity"
	"assignment-2/pkg/errs"
	"assignment-2/repository/order_repository"
	"gorm.io/gorm"
	"log"
)

type orderPg struct {
	db *gorm.DB
}

func NewOrderPG(database *gorm.DB) order_repository.OrderRepository {
	if database == nil {
		log.Panic("Database connection is nil")
	}
	return &orderPg{
		db: database,
	}
}

func (o *orderPg) CreateOrder(orderPayload *entity.Order, itemPayload []entity.Item) (*entity.Order, errs.MessageErr) {
	tx := o.db.Begin()

	if tx.Error != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to start transaction")
	}

	if err := tx.Create(orderPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewBadRequest("Error occurred while trying to create order")
	}

	for _, item := range itemPayload {
		itemCopy := item
		if err := tx.Model(orderPayload).Association("Items").Append(&itemCopy); err != nil {
			tx.Rollback()
			return nil, errs.NewInternalServerError("Error occurred while trying to add item to order")
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Error occurred while trying to commit data to database")
	}

	return orderPayload, nil
}

func (o *orderPg) GetAllData() ([]entity.Order, errs.MessageErr) {
	var allOrder []entity.Order

	if err := o.db.Preload("Items").Find(&allOrder).Error; err != nil {
		return nil, errs.NewNotFound("Error occurred while trying to get all order data")
	}

	return allOrder, nil
}

func (o *orderPg) UpdateOrderById(orderID uint, orderPayload *entity.Order, itemPayload []entity.Item) (*entity.Order, errs.MessageErr) {
	tx := o.db.Begin()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to start transaction")
	}

	order := &entity.Order{}
	if err := tx.First(order, orderID).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewNotFound("Error occurred while trying to find order id")
	}

	order.OrderedAt = orderPayload.OrderedAt
	order.CustomerName = orderPayload.CustomerName

	if err := tx.Save(order).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Error occurred while trying to update")
	}

	if err := tx.Model(order).Association("Items").Clear(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Error occurred while trying to update")
	}

	for _, item := range itemPayload {
		itemCopy := item
		if err := tx.Model(itemPayload).Association("Items").Append(&itemCopy); err != nil {
			tx.Rollback()
			return nil, errs.NewInternalServerError("Error occurred while trying to add new item")
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("Error occurred while trying to commit update")
	}

	return order, nil
}

func (o *orderPg) DeleteOrderById(orderID uint) errs.MessageErr {
	tx := o.db.Begin()

	if err := tx.Error; err != nil {
		return errs.NewInternalServerError("Error occurred while trying to start transaction")
	}

	order := &entity.Order{}
	if err := tx.First(order, orderID).Error; err != nil {
		tx.Rollback()
		return errs.NewNotFound("Error occurred while trying to find order")
	}

	if err := tx.Delete(&order).Error; err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("Error occurred while trying to delete order")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("Error occurred while trying to commit delete")
	}

	return nil
}
