package routes

import (
	"fantasy_pokedex/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/pokemon", controllers.GetPokemon)
	r.GET("/pokemon/:name", controllers.GetPokemonByName)
	r.POST("/pokemon", controllers.CreatePokemon)
	r.DELETE("/pokemon/:name", controllers.DeletePokemon)

	// Admin endpoints
	r.GET("/admin/password-set", controllers.AdminPasswordSetStatus)
	r.POST("/admin/set-password", controllers.AdminSetPassword)
	r.POST("/admin/login", controllers.AdminLogin)

}
