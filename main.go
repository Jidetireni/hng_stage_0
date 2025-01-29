package main

import (
	"hng_stage_0/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	route.RegisterRoutes(server)

	server.Run(":8080")
}
