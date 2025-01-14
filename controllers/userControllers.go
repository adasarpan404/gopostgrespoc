package controllers

import (
	"net/http"

	"github.com/adasarpan404/gopostgrespoc/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (h *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := h.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserController) GetUsers(c *gin.Context) {
	var users []models.User

	if result := h.DB.Preload("Profile").Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if result := h.DB.Preload("Profile").First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
