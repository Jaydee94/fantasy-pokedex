package config

import (
	"log"
	"os"
	"time"

	"fantasy_pokedex/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	var db *gorm.DB
	start := time.Now()
	for {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		if time.Since(start) > time.Minute {
			log.Fatalf("Could not connect to database after 1 minute: %v", err)
		}
		// log.Printf("Waiting for database... (%v)", err)
		time.Sleep(2 * time.Second)
	}

	db.AutoMigrate(&models.Pokemon{})

	DB = db
	// fmt.Println("Database connection established")
}
