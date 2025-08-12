package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/marvini86/car-parts-shop-service/internal/dto"
	"github.com/marvini86/car-parts-shop-service/internal/service"
	"net/http"
	"strconv"
)

// OrderHandler represents a handler for orders
type OrderHandler struct {
	orderService service.OrderService
}

// NewOrderHandler creates a new OrderHandler instance
func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// InitRoutes initializes the routes for the handler
func (h *OrderHandler) InitRoutes(r *gin.RouterGroup) {
	api := r.Group("/orders")
	{
		api.GET("/:id", h.GetOrderByID)
		api.POST("", h.CreateOrder)
	}
}

// GetOrderByID gets an order by ID
// @Summary     Get order by ID
// @Description Get order by ID
// @Tags        Orders
// @Produce     json
// @Param       id  path     int     true  "Order ID"
// @Success     200 {object} dto.OrderDto
// @Router      /orders/{id} [get]
// @Failure     500 {object} dto.ErrorResponseDto
// @Failure     404 {object} dto.ErrorResponseDto
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	ctx := c.Request.Context()

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: err.Error()})
		return
	}

	order, err := h.orderService.GetOrderByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponseDto{Message: err.Error()})
		return
	}

	if order.ID == 0 {
		c.JSON(http.StatusNotFound, dto.ErrorResponseDto{Message: "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CreateOrder creates a new order
// @Summary     Create order
// @Description Create order
// @Tags        Orders
// @Produce     json
// @Param       orderCreate  body     dto.OrderRequest  true  "Order request"
// @Success     201 {object} dto.OrderDto
// @Router      /orders [post]
// @Failure     400 {object} dto.ErrorResponseDto
// @Failure     500 {object} dto.ErrorResponseDto
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var orderCreate dto.OrderRequest
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&orderCreate); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: err.Error()})
		return
	}

	order, err := h.orderService.CreateOrder(ctx, orderCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponseDto{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
	return
}
