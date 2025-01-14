package main

import (
	"log"

	"github.com/adasarpan404/gopostgrespoc/config"
	"github.com/adasarpan404/gopostgrespoc/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()

	router := gin.Default()

	routes.SetupUserRoutes(router, db)

	log.Println("Server is running on http://localhost:8080")
	router.Run(":8080")
}
