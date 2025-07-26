package dto

// ItemDto represents a single item
type ItemDto struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	AvailableQuantity int     `json:"availableQuantity"`
	Image             string  `json:"image"`
	Category          string  `json:"category"`
}
