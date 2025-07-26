package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marvini86/car-parts-shop-service/internal/handler"
	"github.com/marvini86/car-parts-shop-service/internal/service"
	"gorm.io/gorm"
	"os"
)

// ServerConfig represents the configuration for the server
type ServerConfig struct {
	db  *gorm.DB
	gin *gin.Engine
}

// NewServerConfig creates a new ServerConfig instance
func NewServerConfig(db *gorm.DB) *ServerConfig {
	return &ServerConfig{
		db: db,
	}
}

var apiPort string

// Init initializes the server
func (s *ServerConfig) Init() {
	if port, ok := os.LookupEnv("API_PORT"); ok {
		apiPort = port
	}
	s.gin = gin.Default()

	s.initRoutes()

	s.gin.Run(fmt.Sprintf(":%s", apiPort))
}

// initRoutes initializes the routes for the server
func (s *ServerConfig) initRoutes() {
	r := s.gin.Group("/api/v1")

	handler.NewItemHandler(service.NewItemService(s.db)).InitRoutes(r)
	handler.NewOrderHandler(service.NewOrderService(s.db)).InitRoutes(r)
}
