package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/marvini86/car-parts-shop-service/internal/dto"
	"github.com/marvini86/car-parts-shop-service/internal/entity"
	"github.com/marvini86/car-parts-shop-service/internal/grpc/inventory"
	"gorm.io/gorm"
	"log"
)

// OrderService provides methods for interacting with orders
type OrderService interface {
	CreateOrder(ctx context.Context, orderCreate dto.OrderRequest) (createdOrder dto.SlimOrderDto, err error)
	GetOrderByID(ctx context.Context, id int) (order dto.OrderDto, err error)
}

// orderService implements OrderService
type orderService struct {
	db *gorm.DB
}

// NewOrderService creates a new OrderService instance
func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{db: db}
}

// GetOrderByID gets an order by ID
func (s *orderService) GetOrderByID(ctx context.Context, id int) (order dto.OrderDto, err error) {
	var orderEntity entity.Order
	err = s.db.WithContext(ctx).Preload("User").Preload("Items.Item").Preload("PaymentDetails").Preload("DeliveryAddress").First(&orderEntity, id).Error
	copier.Copy(&order, &orderEntity)

	items := make([]dto.OrderItemDto, 0, len(orderEntity.Items))

	for _, item := range orderEntity.Items {
		var itemDto dto.OrderItemDto
		copier.Copy(&itemDto, &item)
		itemDto.Name = item.Item.Name
		items = append(items, itemDto)
	}
	order.Items = items

	return
}

// CreateOrder creates a new order
func (s *orderService) CreateOrder(ctx context.Context, orderCreate dto.OrderRequest) (createdOrder dto.SlimOrderDto, err error) {
	s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var order = entity.Order{
			UserID:     orderCreate.UserID,
			Status:     entity.Pending,
			TotalValue: orderCreate.TotalValue,
		}

		for _, item := range orderCreate.Items {
			var itemEntity entity.Item
			err = tx.WithContext(ctx).First(&itemEntity, item.ItemID).Error
			if err != nil {
				log.Printf("Error getting item with ID %d: %v", item.ItemID, err)
				return fmt.Errorf("failed to get item with ID %d: %w", item.ItemID, err)
			}
			log.Printf("Item found with ID %d and & code %s", item.ItemID, itemEntity.CodeIntegration)
			res, err := inventory.CheckAvailability(context.Background(), itemEntity.CodeIntegration)
			if err != nil {
				log.Printf("Error checking availability for item with ID %d: %v", item.ItemID, err)
				//return err
			}
			log.Printf("Availability for item with ID %d: %v", item.ItemID, res)
		}

		// Save Order
		if err := tx.WithContext(ctx).Create(&order).Error; err != nil {
			return fmt.Errorf("failed to create order: %w", err)
		}

		// Set foreign keys for related tables
		for _, item := range orderCreate.Items {
			var itemEntity entity.OrderItem
			itemEntity.ItemID = item.ItemID
			itemEntity.Quantity = item.Quantity
			itemEntity.Price = item.Price
			itemEntity.OrderID = order.ID
			if err := tx.WithContext(ctx).Create(&itemEntity).Error; err != nil {
				return fmt.Errorf("failed to create order item: %w", err)
			}
		}

		var payment entity.OrderPaymentDetails
		payment.OrderID = order.ID
		payment.CardNumber = orderCreate.PaymentDetails.CardNumber
		payment.ExpiryDate = orderCreate.PaymentDetails.ExpiryDate
		payment.CVV = orderCreate.PaymentDetails.CVV

		// Save Payment Details
		if err := tx.WithContext(ctx).Create(&payment).Error; err != nil {
			return fmt.Errorf("failed to create payment details: %w", err)
		}

		var address entity.OrderDeliveryAddress
		address.OrderID = order.ID
		address.Address = orderCreate.DeliveryAddress.Address
		address.City = orderCreate.DeliveryAddress.City
		address.State = orderCreate.DeliveryAddress.State
		address.ZipCode = orderCreate.DeliveryAddress.ZipCode
		address.Country = orderCreate.DeliveryAddress.Country

		// Save Delivery Address
		if err := tx.WithContext(ctx).Create(&address).Error; err != nil {
			return fmt.Errorf("failed to create delivery address: %w", err)
		}

		copier.Copy(&createdOrder, &order)

		return nil
	})

	return
}
