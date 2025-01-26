package models

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"example.com/movies-api/db"
	"example.com/movies-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string
	Password string
}

func (user User) CreateUser() error {
	collection := db.MongoClient.Database(db.Db).Collection("users")
	// Debug: Print the user being created
	fmt.Printf("DEBUG: Creating user - Email: %s, Password: %s\n", user.Email, user.Password)
	hashpassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashpassword)
	// inserting new user
	inserted, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	// Debug: Print the inserted user ID
	fmt.Printf("DEBUG: User created successfully - ID: %v\n", inserted.InsertedID)
	user.ID = inserted.InsertedID.(primitive.ObjectID)
	return nil
}

// validate a user
func (user *User) Validate() error {
	collection := db.MongoClient.Database(db.Db).Collection("users")
	filter := bson.M{"email": user.Email}
	// Trim leading/trailing spaces from the email
	user.Email = strings.TrimSpace(user.Email)

	// Debug: Print the email being queried
	fmt.Printf("DEBUG: Querying user with email: %s\n", user.Email)
	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("DEBUG: User not found")
			return errors.New("credential invalid")
		}
		return err
	}
	// Assign the retrieved ID to the user struct
	user.ID = result.ID
	// Debug: Print the email and hashed password retrieved from the database
	fmt.Printf("DEBUG: Found user - Email: %s, Hashed Password: %s\n", user.Email, user.Password)
	passwordIsValid := utils.CheckPasswordHash(user.Password, result.Password)
	if !passwordIsValid {
		fmt.Println("DEBUG: Password mismatch")
		return errors.New("credential invalid")
	}
	return nil
}
