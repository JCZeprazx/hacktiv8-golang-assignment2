package handler

import (
	"assignment-2/database"
	"assignment-2/dto"
	"assignment-2/repository/order_repository/order_postgres"
	"assignment-2/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func StartApp() {
	db := database.GetDataBaseInstance()

	orderRepo := order_postgres.NewOrderPG(db)
	orderService := service.NewOrderService(orderRepo)

	r := gin.Default()

	r.POST("/order", func(ctx *gin.Context) {
		var request dto.OrderRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		order, err := orderService.CreateOrder(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(order.StatusCode, order)
	})

	r.GET("/all-orders", func(ctx *gin.Context) {
		allOrders, err := orderService.GetAllData()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(allOrders.StatusCode, allOrders)
	})

	r.PATCH("/order/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		orderID, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		var request dto.OrderRequest
		updateOrder, err := orderService.UpdateOrderById(uint(orderID), request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(updateOrder.StatusCode, updateOrder)
	})

	r.DELETE("order/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		orderID, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		deleteOrder, err := orderService.DeleteOrderById(uint(orderID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(deleteOrder.StatusCode, deleteOrder)
	})

	if err := r.Run(":8080"); err != nil {
		return
	}
}
