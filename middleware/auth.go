package middleware

import (
	"net/http"
	"strings"

	"example.com/movies-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	//Get the authorization header
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token format"})
		return
	}
	userId, err := utils.VerifyToken(token) //verifying if the token is valid
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}
	context.Set("userId", userId)
}
