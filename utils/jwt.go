package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId primitive.ObjectID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

// verify token
func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpexted signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", errors.New("could not parse token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return "", errors.New("Invalid token!")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("invalid user ID in token")
	}
	return userId, nil
}
