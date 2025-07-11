package controllers

import (
	"encoding/base64"
	"fantasy_pokedex/config"
	"fantasy_pokedex/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreatePokemonRequest struct {
	Name         string   `json:"Name"`
	Types        []string `json:"Types"`
	PokedexEntry string   `json:"PokedexEntry"`
	ImageBase64  string   `json:"ImageBase64"` // base64 String vom Frontend
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
	var req CreatePokemonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Base64 String bereinigen (falls mit Data-URL-Schema)
	base64Data := req.ImageBase64
	if idx := strings.Index(base64Data, ","); idx != -1 {
		base64Data = base64Data[idx+1:]
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 image data"})
		return
	}

	pokemon := models.Pokemon{
		Name:         req.Name,
		Types:        req.Types,
		PokedexEntry: req.PokedexEntry,
		ImageData:    imageData,
		Appearance:   req.Appearance,
		Attacks:      req.Attacks,
		Ability:      req.Ability,
		HeightM:      req.HeightM,
		WeightKg:     req.WeightKg,
		Category:     req.Category,
	}

	if err := config.DB.Create(&pokemon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create pokemon"})
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
