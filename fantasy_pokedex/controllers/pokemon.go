package controllers

import (
	"encoding/base64"
	"fantasy_pokedex/config"
	"fantasy_pokedex/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePokemonInput struct {
	Name         string   `json:"Name"`
	Types        []string `json:"Types"`
	PokedexEntry string   `json:"PokedexEntry"`
	ImageData    string   `json:"ImageData"`
	Appearance   string   `json:"Appearance"`
	Attacks      []string `json:"Attacks"`
	Ability      string   `json:"Ability"`
	HeightM      float64  `json:"HeightM"`
	WeightKg     float64  `json:"WeightKg"`
	Category     string   `json:"Category"`
}

type PokemonResponse struct {
	ID           uint     `json:"ID"`
	Name         string   `json:"Name"`
	Types        []string `json:"Types"`
	PokedexEntry string   `json:"PokedexEntry"`
	ImageBase64  string   `json:"ImageBase64"`
	Appearance   string   `json:"Appearance"`
	Attacks      []string `json:"Attacks"`
	Ability      string   `json:"Ability"`
	HeightM      float64  `json:"HeightM"`
	WeightKg     float64  `json:"WeightKg"`
	Category     string   `json:"Category"`
}

func CreatePokemon(c *gin.Context) {
	var input CreatePokemonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 📥 Base64 dekodieren
	imageBytes, err := base64.StdEncoding.DecodeString(input.ImageData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bild konnte nicht dekodiert werden"})
		return
	}

	pokemon := models.Pokemon{
		Name:         input.Name,
		Types:        input.Types,
		PokedexEntry: input.PokedexEntry,
		ImageData:    imageBytes,
		Appearance:   input.Appearance,
		Attacks:      input.Attacks,
		Ability:      input.Ability,
		HeightM:      input.HeightM,
		WeightKg:     input.WeightKg,
		Category:     input.Category,
	}

	if err := config.DB.Create(&pokemon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Speichern"})
		return
	}

	c.JSON(http.StatusCreated, pokemon)
}

func toPokemonResponse(p models.Pokemon) PokemonResponse {
	return PokemonResponse{
		ID:           p.ID,
		Name:         p.Name,
		Types:        p.Types,
		PokedexEntry: p.PokedexEntry,
		ImageBase64:  base64.StdEncoding.EncodeToString(p.ImageData),
		Appearance:   p.Appearance,
		Attacks:      p.Attacks,
		Ability:      p.Ability,
		HeightM:      p.HeightM,
		WeightKg:     p.WeightKg,
		Category:     p.Category,
	}
}

func GetPokemon(c *gin.Context) {
	var pokemons []models.Pokemon
	config.DB.Find(&pokemons)

	var response []PokemonResponse
	for _, p := range pokemons {
		response = append(response, toPokemonResponse(p))
	}

	c.JSON(http.StatusOK, response)
}

func GetPokemonByID(c *gin.Context) {
	var pokemon models.Pokemon
	id := c.Param("id")

	if err := config.DB.First(&pokemon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}

	c.JSON(http.StatusOK, toPokemonResponse(pokemon))
}
