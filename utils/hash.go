package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

func CheckPasswordHash(password, hashedPassword string) bool {
	fmt.Printf("DEBUG: Comparing password: %s with hash: %s\n", password, hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Printf("DEBUG: Password comparison error: %v\n", err)
		return false
	}
	return true
}
