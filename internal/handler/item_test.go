package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/marvini86/car-parts-shop-service/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockItemService struct {
	mock.Mock
}

func (m *MockItemService) GetAllItems(ctx context.Context) ([]dto.ItemDto, error) {
	args := m.Called()
	return args.Get(0).([]dto.ItemDto), args.Error(1)
}

func TestGetAllItems_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockItemService)
	handler := NewItemHandler(mockService)

	// Prepare mock data
	mockItems := []dto.ItemDto{
		{ID: 1, Name: "Brake Pad", Price: 49.99, Category: "Brakes"},
	}

	mockService.On("GetAllItems").Return(mockItems, nil)

	// Setup Gin router with the handler
	r := gin.Default()
	r.GET("/items", handler.GetAllItems)

	req, _ := http.NewRequest(http.MethodGet, "/items", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)

	// Optionally check response body content as JSON
	// But here we just ensure status and call
}

func TestGetAllItems_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockItemService)
	handler := NewItemHandler(mockService)

	mockService.On("GetAllItems").Return(nil, assert.AnError)

	r := gin.Default()
	r.GET("/items", handler.GetAllItems)

	req, _ := http.NewRequest(http.MethodGet, "/items", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockService.AssertExpectations(t)
}
