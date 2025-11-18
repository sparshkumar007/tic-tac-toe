package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	supabasePassword := os.Getenv("SUPABASE_DB_PASSWORD")
	databaseUrlTemplate := os.Getenv("DATABASE_URL")
	dsn := fmt.Sprintf(
		databaseUrlTemplate,
		supabasePassword,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	DB = db
	fmt.Println("Connected to Supabase successfully!")
}
