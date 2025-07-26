package dto

import "time"

type SlimOrderDto struct {
	ID         int       `json:"id"`
	Status     string    `json:"status"`
	TotalValue float64   `json:"totalValue"`
	CreatedAt  time.Time `json:"createdAt"`
}
type OrderDto struct {
	ID              int                     `json:"id"`
	User            UserDto                 `json:"user"`
	Status          string                  `json:"status"`
	TotalValue      float64                 `json:"totalValue"`
	CreatedAt       string                  `json:"createdAt"`
	Items           []OrderItemDto          `json:"items"`
	PaymentDetails  OrderPaymentDetailsDto  `json:"paymentDetails"`
	DeliveryAddress OrderDeliveryAddressDto `json:"deliveryAddress"`
}

type OrderItemDto struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type OrderPaymentDetailsDto struct {
	CardNumber string `json:"cardNumber"`
	ExpiryDate string `json:"expiryDate"`
	CVV        string `json:"cvv"`
}

type OrderDeliveryAddressDto struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
	Country string `json:"country"`
}
