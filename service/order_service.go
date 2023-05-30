package service

import (
	"assignment-2/dto"
	"assignment-2/entity"
	"assignment-2/pkg/errs"
	"assignment-2/repository/order_repository"
	"net/http"
)

type orderService struct {
	orderRepo order_repository.OrderRepository
}

type OrderService interface {
	CreateOrder(payload dto.OrderRequest) (*dto.NewOrderResponse, errs.MessageErr)
	GetAllData() (*dto.GetAllOrderResponse, errs.MessageErr)
	UpdateOrderById(orderID uint, payload dto.OrderRequest) (*dto.UpdateOrderById, errs.MessageErr)
	DeleteOrderById(orderID uint) (*dto.DeleteOrderById, errs.MessageErr)
}

func NewOrderService(orderRepo order_repository.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) CreateOrder(payload dto.OrderRequest) (*dto.NewOrderResponse, errs.MessageErr) {
	orderPayload := payload.OrderRequestToEntity()

	itemsPayload := []entity.Item{}
	for _, items := range payload.Items {
		item := items.ItemRequestToEntity()
		itemsPayload = append(itemsPayload, *item)
	}

	orderRequest, err := o.orderRepo.CreateOrder(orderPayload, itemsPayload)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to create")
	}

	response := &dto.NewOrderResponse{
		StatusCode: http.StatusCreated,
		Message:    "Success",
		Data: dto.OrderRequest{
			OrderedAt:    orderRequest.OrderedAt,
			CustomerName: orderRequest.CustomerName,
			Items:        payload.Items,
		},
	}

	return response, nil
}

func (o *orderService) GetAllData() (*dto.GetAllOrderResponse, errs.MessageErr) {
	allOrders, err := o.orderRepo.GetAllData()

	if err != nil {
		return nil, errs.NewBadRequest("Cannot found data")
	}

	orderData := []dto.OrderResponse{}
	for _, order := range allOrders {
		items := []dto.ItemResponse{}
		for _, item := range order.Items {
			itemResponse := dto.ItemResponse{
				ID:          item.ID,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
				OrderID:     item.OrderID,
			}
			items = append(items, itemResponse)
		}
		orderResponse := dto.OrderResponse{
			ID:           order.ID,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			CustomerName: order.CustomerName,
			Items:        items,
		}
		orderData = append(orderData, orderResponse)
	}
	response := &dto.GetAllOrderResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       orderData,
	}

	return response, nil
}

func (o *orderService) UpdateOrderById(orderID uint, payload dto.OrderRequest) (*dto.UpdateOrderById, errs.MessageErr) {
	orderPayload := payload.OrderRequestToEntity()

	itemsPayload := []entity.Item{}
	for _, item := range payload.Items {
		items := item.ItemRequestToEntity()
		itemsPayload = append(itemsPayload, *items)
	}

	updateOrder, err := o.orderRepo.UpdateOrderById(orderID, orderPayload, itemsPayload)
	if err != nil {
		return nil, errs.NewBadRequest("Cannot find order to process update")
	}

	itemResponse := []dto.ItemResponse{}
	for _, item := range updateOrder.Items {
		item := dto.ItemResponse{
			ID:          item.ID,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     item.OrderID,
		}
		itemResponse = append(itemResponse, item)
	}

	response := &dto.UpdateOrderById{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data: dto.OrderResponse{
			ID:           updateOrder.ID,
			CreatedAt:    updateOrder.CreatedAt,
			UpdatedAt:    updateOrder.UpdatedAt,
			CustomerName: updateOrder.CustomerName,
			Items:        itemResponse,
		},
	}

	return response, nil
}

func (o *orderService) DeleteOrderById(orderID uint) (*dto.DeleteOrderById, errs.MessageErr) {
	if err := o.orderRepo.DeleteOrderById(orderID); err != nil {
		return nil, errs.NewNotFound("Error occurred while trying to find order")
	}

	deleteResponse := &dto.DeleteOrderById{
		StatusCode: http.StatusOK,
		Message:    "Success delete",
	}

	return deleteResponse, nil
}
