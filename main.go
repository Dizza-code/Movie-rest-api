package main

import (
	"log"

	"example.com/movies-api/db"
	"example.com/movies-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Init session store
	server.Use(gin.Logger())

	db.ConnectDatabase()
	routes.RegisterRoutes(server)

	log.Println("Server started!")
	server.Run() // Default Port 8080
}
