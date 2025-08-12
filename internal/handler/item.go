package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/marvini86/car-parts-shop-service/internal/dto"
	"github.com/marvini86/car-parts-shop-service/internal/service"
	"net/http"
)

// ItemHandler represents a handler for items
type ItemHandler struct {
	itemService service.ItemService
}

// NewItemHandler creates a new ItemHandler instance
func NewItemHandler(itemService service.ItemService) *ItemHandler {
	return &ItemHandler{
		itemService: itemService,
	}
}

// InitRoutes initializes the routes for the handler
func (h *ItemHandler) InitRoutes(r *gin.RouterGroup) {
	api := r.Group("/items")
	{
		api.GET("", h.GetAllItems)
	}
}

// GetAllItems gets all items
// @Summary     Get all items
// @Description Get all available items
// @Tags        Items
// @Produce     json
// @Success     200 {object} []dto.ItemDto
// @Router      /items [get]
// @Failure     500 {object} dto.ErrorResponseDto
func (h *ItemHandler) GetAllItems(c *gin.Context) {
	ctx := c.Request.Context()

	items, err := h.itemService.GetAllItems(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponseDto{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
	return
}
