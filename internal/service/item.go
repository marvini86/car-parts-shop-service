package service

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/marvini86/car-parts-shop-service/internal/dto"
	"github.com/marvini86/car-parts-shop-service/internal/entity"
	"gorm.io/gorm"
)

// ItemService provides methods for interacting with items
type ItemService interface {
	GetAllItems(ctx context.Context) (items []dto.ItemDto, err error)
}

// itemService implements ItemService
type itemService struct {
	db *gorm.DB
}

// NewItemService creates a new ItemService instance
func NewItemService(db *gorm.DB) ItemService {
	return &itemService{
		db: db,
	}
}

// GetAllItems returns all items
func (s *itemService) GetAllItems(ctx context.Context) (items []dto.ItemDto, err error) {
	var itemsEntity []entity.Item
	err = s.db.WithContext(ctx).Preload("Category").Find(&itemsEntity).Error
	if err != nil {
		return
	}
	items = s.toDtos(itemsEntity)
	return
}

// toDtos converts entity.Item to dto.ItemDto
func (s *itemService) toDtos(items []entity.Item) []dto.ItemDto {
	dtos := make([]dto.ItemDto, 0, len(items))
	for _, item := range items {
		var dto dto.ItemDto
		copier.Copy(&dto, &item)
		dto.Category = item.Category.Name
		dtos = append(dtos, dto)
	}
	return dtos
}
