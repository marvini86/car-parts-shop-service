package entity

import "time"

type OrderStatus string

const (
	Pending   OrderStatus = "PENDING"
	Shipped   OrderStatus = "SHIPPED"
	Delivered OrderStatus = "DELIVERED"
)

type Order struct {
	ID              int
	UserID          int
	User            User
	TotalValue      float64
	Status          OrderStatus
	Items           []OrderItem
	PaymentDetails  OrderPaymentDetails
	DeliveryAddress OrderDeliveryAddress
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type OrderItem struct {
	ID        int
	Item      Item
	OrderID   int
	ItemID    int
	Quantity  int
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderPaymentDetails struct {
	ID         int
	OrderID    int
	CardNumber string
	ExpiryDate string
	CVV        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OrderDeliveryAddress struct {
	ID        int
	OrderID   int
	Address   string
	City      string
	State     string
	ZipCode   string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
