package main

import (
	"fantasy_pokedex/config"
	"fantasy_pokedex/models"
	"fantasy_pokedex/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func createDummyPokemon() {
	// Bild lokal laden (z.B. ./assets/testmon.png)
	imagePath := "./assets/dummy.png"
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		log.Printf("Fehler beim Lesen des Bildes: %v", err)
		return
	}

	dummyPokemon := models.Pokemon{
		ID:           1,
		Name:         "Testmon",
		Types:        []string{"Fire", "Flying"},
		PokedexEntry: "Testmon is a fictional Pokemon used for testing purposes.",
		ImageData:    imageData, // Bildbytes hier rein
		Appearance:   "A small, red, dragon-like creature with wings.",
		Attacks:      []string{"Flame Burst", "Wing Attack"},
		Ability:      "Test Ability",
		HeightM:      1.0,
		WeightKg:     10.0,
		Category:     "Test",
	}

	result := config.DB.Create(&dummyPokemon)
	if result.Error != nil {
		log.Printf("Error creating dummy Pokemon: %v", result.Error)
	} else {
		log.Printf("Created dummy Pokemon: %s", dummyPokemon.Name)
	}
}

func main() {
	config.ConnectDatabase()

	if gin.Mode() != gin.ReleaseMode {
		createDummyPokemon()
	}

	r := gin.Default()

	// 👉 CORS konfigurieren
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(r)

	r.Run() // default: :8080
}
