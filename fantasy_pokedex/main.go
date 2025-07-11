package main

import (
	"fantasy_pokedex/config"
	"fantasy_pokedex/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run() // default: localhost:8080
}
