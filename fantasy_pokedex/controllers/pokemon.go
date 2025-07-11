package controllers

import (
	"fantasy_pokedex/config"
	"fantasy_pokedex/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPokemon(c *gin.Context) {
	var pokemons []models.Pokemon
	config.DB.Find(&pokemons)
	c.JSON(http.StatusOK, pokemons)
}

func GetPokemonByID(c *gin.Context) {
	var pokemon models.Pokemon
	id := c.Param("id")

	if err := config.DB.First(&pokemon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}

	c.JSON(http.StatusOK, pokemon)
}

func CreatePokemon(c *gin.Context) {
	var pokemon models.Pokemon
	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&pokemon)
	c.JSON(http.StatusCreated, pokemon)
}
