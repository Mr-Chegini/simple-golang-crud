package database

import (
	"log"

	"github.com/Mr-Chegini/simple-golang-crud/models"
)

// RunMigrations creates the necessary tables using GORM
func RunMigrations() error {
	// Auto migrate creates tables from models
	if err := DB.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		return err
	}

	log.Println("✅ Database migrations completed successfully")
	return nil
}
