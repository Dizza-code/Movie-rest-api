package controllers

import (
	"fmt"
	"net/http"

	"example.com/movies-api/models"
	"example.com/movies-api/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
	err = user.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user successfully created"})
}
func LoginUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.Validate()
	if err != nil {
		fmt.Println("Error parsing request:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
