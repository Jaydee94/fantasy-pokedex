package routes

import (
	"fantasy_pokedex/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/pokemon", controllers.GetPokemon)
	r.GET("/pokemon/:id", controllers.GetPokemonByID)
	r.POST("/pokemon", controllers.CreatePokemon)

}
