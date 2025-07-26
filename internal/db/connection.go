package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// OpenConnection opens a database connection
func OpenConnection() (db *gorm.DB) {
	sqlConn := fmt.Sprintf("host=%s	port=%s	user=%s	password=%s	dbname=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB"))
	db, err := gorm.Open(postgres.Open(sqlConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return
}
