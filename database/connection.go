package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database instance
var DB *gorm.DB

// InitDB initializes the database connection using GORM + PostgreSQL
func InitDB() error {
	var err error

	// Get database connection string from environment or use default
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Default connection string (development)
		dsn = "host=localhost user=postgres password=postgres dbname=simple_crud port=5432 sslmode=disable"
	}

	// Connect to PostgreSQL database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("✅ Database connection established successfully")
	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}
	return sqlDB.Close()
}

// GetDB returns the global database instance
func GetDB() *gorm.DB {
	return DB
}
