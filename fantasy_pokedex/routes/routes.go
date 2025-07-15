package routes

import (
	"fantasy_pokedex/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/pokemon", controllers.GetPokemon)
	r.GET("/pokemon/:name", controllers.GetPokemonByName)
	r.POST("/pokemon", controllers.CreatePokemon)

}
