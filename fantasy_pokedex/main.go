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
	imagePath := "./assets/testly.png"
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		log.Printf("Error reading image: %v", err)
		return
	}

	dummyPokemon := models.Pokemon{
		Name:         "Testly",
		Types:        []string{"Electric", "Psychic"},
		PokedexEntry: "Testly is known for instantly recognizing weaknesses and patterns in any environment. It loves to test new things – whether machines, berries, or even trainers. It is said that it can see through its opponent's strategy just by looking.",
		ImageData:    imageData,
		Appearance:   "A small, energetic Pokémon with technical details, glowing eyes, and a small screen on its forehead.",
		Attacks: []string{
			"Tackle",
			"Spark",
			"Scan Gaze",
			"Psychokinesis",
			"Trickery",
			"Test Shock",
		},
		Ability:   "Analyzer (boosts moves when Testly moves last); Hidden Ability: Data Surge (raises Special Attack with each defeated opponent)",
		HeightM:   0.8,
		WeightKg:  12.5,
		Category:  "Analysis Pokémon",
		HP:        80,
		Attack:    90,
		Defense:   70,
		Speed:     100,
		SpAttack:  110,
		SpDefense: 85,
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(r)

	// Health endpoint for Kubernetes
	r.GET("/healthz", func(c *gin.Context) {
		sqlDB, err := config.DB.DB()
		if err != nil {
			c.JSON(500, gin.H{"status": "db error", "error": err.Error()})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{"status": "db not ready", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run() // default: :8080
}
