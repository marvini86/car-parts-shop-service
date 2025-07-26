package dto

type OrderRequest struct {
	Items           []OrderItemRequest          `json:"items"`
	TotalValue      float64                     `json:"totalValue"`
	UserID          int                         `json:"userId"`
	PaymentDetails  OrderPaymentDetailsRequest  `json:"paymentDetails"`
	DeliveryAddress OrderDeliveryAddressRequest `json:"deliveryAddress"`
}

type OrderItemRequest struct {
	ItemID   int     `json:"itemId"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type OrderPaymentDetailsRequest struct {
	CardNumber string `json:"cardNumber"`
	ExpiryDate string `json:"expiryDate"`
	CVV        string `json:"cvv"`
}

type OrderDeliveryAddressRequest struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
	Country string `json:"country"`
}
