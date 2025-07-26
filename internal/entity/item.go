package entity

import "time"

type Item struct {
	ID                int
	CodeIntegration   string
	Name              string
	Description       string
	Price             float64
	AvailableQuantity int
	Image             string
	Category          ItemCategory
	CategoryID        int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type ItemCategory struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
