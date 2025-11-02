package db

import (
	"log"
	"os"

	"order-service/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB initializes and returns GORM DB connection
func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Auto-migrate models
	err = db.AutoMigrate(&model.Order{}, &model.OrderItem{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	log.Println("PostgreSQL connected and migrated")
	return db
}