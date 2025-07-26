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
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	ctx := c.Request.Context()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.GetOrderByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if order.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CreateOrder creates a new order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var orderCreate dto.OrderRequest
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&orderCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.CreateOrder(ctx, orderCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
	return
}
