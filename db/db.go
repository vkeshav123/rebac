package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rebac/models"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Example fallback DSN, update as needed
		dsn = "host=localhost user=postgres password=postgres dbname=rebac port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	DB = db

	// Auto-migrate all models
	err = db.AutoMigrate(
		&models.Model{},
		&models.Type{},
		&models.Relation{},
		&models.DirectRelation{},
		&models.ImpliedRelation{},
		&models.Policy{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Println("Database connection established and migrated")
}
