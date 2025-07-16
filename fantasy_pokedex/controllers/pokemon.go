package controllers

import (
	"encoding/base64"
	"fantasy_pokedex/config"
	"fantasy_pokedex/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeletePokemon deletes a Pokemon by name
func DeletePokemon(c *gin.Context) {
	name := c.Param("name")
	var pokemon models.Pokemon
	if err := config.DB.First(&pokemon, "name = ?", name).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}
	if err := config.DB.Delete(&pokemon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Pokemon"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pokemon deleted successfully"})
}

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
	HP           int      `json:"HP"`
	Attack       int      `json:"Attack"`
	Defense      int      `json:"Defense"`
	Speed        int      `json:"Speed"`
	SpAttack     int      `json:"SpAttack"`
	SpDefense    int      `json:"SpDefense"`
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
	HP           int      `json:"HP"`
	Attack       int      `json:"Attack"`
	Defense      int      `json:"Defense"`
	Speed        int      `json:"Speed"`
	SpAttack     int      `json:"SpAttack"`
	SpDefense    int      `json:"SpDefense"`
}

func CreatePokemon(c *gin.Context) {
	var input CreatePokemonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageBytes, err := base64.StdEncoding.DecodeString(input.ImageData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode image"})
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
		HP:           input.HP,
		Attack:       input.Attack,
		Defense:      input.Defense,
		Speed:        input.Speed,
		SpAttack:     input.SpAttack,
		SpDefense:    input.SpDefense,
	}

	if err := config.DB.Create(&pokemon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save Pokémon"})
		return
	}

	c.JSON(http.StatusCreated, pokemon)
}

func toPokemonResponse(p models.Pokemon) PokemonResponse {
	return PokemonResponse{
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
		HP:           p.HP,
		Attack:       p.Attack,
		Defense:      p.Defense,
		Speed:        p.Speed,
		SpAttack:     p.SpAttack,
		SpDefense:    p.SpDefense,
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

func GetPokemonByName(c *gin.Context) {
	var pokemon models.Pokemon
	name := c.Param("name")

	if err := config.DB.First(&pokemon, name).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}

	c.JSON(http.StatusOK, toPokemonResponse(pokemon))
}
