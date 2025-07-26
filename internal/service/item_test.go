package service_test

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marvini86/car-parts-shop-service/internal/db"
	"github.com/marvini86/car-parts-shop-service/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllItems(t *testing.T) {
	db, mock, teardown := db.NewMockDB(t)
	defer teardown()

	ctx := context.Background()

	// Prepare mock rows
	rows := sqlmock.NewRows([]string{"id", "name", "price", "category_id"}).
		AddRow(1, "Brake Pad", 49.99, 10)

	categoryRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(10, "Brakes")

	// Expect item query
	mock.ExpectQuery(`SELECT .* FROM "items"`).
		WillReturnRows(rows)

	// Expect category preload query
	mock.ExpectQuery(`SELECT .* FROM "item_categories"`).
		WillReturnRows(categoryRows)

	itemService := service.NewItemService(db)

	items, err := itemService.GetAllItems(ctx)
	assert.NoError(t, err)
	assert.Len(t, items, 1)

	assert.Equal(t, "Brake Pad", items[0].Name)
	assert.Equal(t, "Brakes", items[0].Category)
}

func TestGetAllItems_DBError(t *testing.T) {
	db, mock, cleanup := db.NewMockDB(t)
	defer cleanup()
	ctx := context.Background()

	// Simulate a db error when querying items
	mock.ExpectQuery(`SELECT .* FROM "items"`).
		WillReturnError(fmt.Errorf("mocked db failure"))

	itemService := service.NewItemService(db)

	items, err := itemService.GetAllItems(ctx)

	assert.Error(t, err)
	assert.Nil(t, items)
	assert.EqualError(t, err, "mocked db failure")
}
