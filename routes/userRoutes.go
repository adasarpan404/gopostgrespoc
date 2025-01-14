package routes

import (
	"github.com/adasarpan404/gopostgrespoc/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	userHandler := controllers.NewUserHandler(db)
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUserById)
}
