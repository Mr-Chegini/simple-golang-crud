package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the global database instance
var DB *gorm.DB

// InitDB initializes the database connection using GORM + SQLite
func InitDB() error {
	var err error

	// Connect to SQLite database
	DB, err = gorm.Open(sqlite.Open("database/app.db"), &gorm.Config{})
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
