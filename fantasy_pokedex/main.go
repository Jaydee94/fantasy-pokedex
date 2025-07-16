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
		log.Printf("Fehler beim Lesen des Bildes: %v", err)
		return
	}

	dummyPokemon := models.Pokemon{
		Name:         "Testly",
		Types:        []string{"Elektro", "Psycho"},
		PokedexEntry: "Testly ist bekannt dafür, in jeder Umgebung sofort Schwachstellen und Muster zu erkennen. Es liebt es, neue Dinge zu testen – ob Maschinen, Beeren oder sogar Trainer. Man sagt, es kann durch bloßes Anschauen die Strategie seines Gegners durchschauen.",
		ImageData:    imageData,
		Appearance:   "Ein kleines, energiegeladenes Pokémon mit technischen Details, leuchtenden Augen und einem kleinen Bildschirm auf der Stirn.",
		Attacks: []string{
			"Tackle",
			"Ladungsstoß",
			"Scannerblick",
			"Psychokinese",
			"Trickbetrug",
			"Testschock",
		},
		Ability:  "Analyser (verstärkt Attacken, wenn Testly als letztes angreift); Versteckte Fähigkeit: Datenrausch (steigert Spezial-Angriff bei jedem besiegten Gegner)",
		HeightM:  0.8,
		WeightKg: 12.5,
		Category: "Analyse-Pokémon",
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

	r.Run() // default: :8080
}
