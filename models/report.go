package models

import (
	"bytes"
	"encoding/csv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateCSV(movies []bson.M) ([]byte, error) {
	var csvData bytes.Buffer
	writer := csv.NewWriter(&csvData)

	//write the csv header
	header := []string{"Movie", "Actors", "Created At", "User Email"}
	if err := writer.Write(header); err != nil {
		return nil, err
	}
	// write the movie data
	for _, movie := range movies {

		actorInterfaces := movie["actors"].(bson.A) // Get the actors as a bson.A (slice of interface{})
		actors := make([]string, len(actorInterfaces))
		for i, actor := range actorInterfaces {
			actors[i] = actor.(string) // Convert each actor to a string
		}
		actorsString := strings.Join(actors, ", ") // Join the actors into a comma-separated string
		// //convert actors slice to a comma seperated string
		// actors := strings.Join(movie["actors"].(bson.A).ToStringSlice(), ", ")
		// Convert created_at to a string
		createdAt := movie["created_at"].(primitive.DateTime).Time().Format(time.RFC3339)

		//write the row
		row := []string{
			movie["movie"].(string),
			actorsString,
			createdAt,
			movie["user_email"].(string),
		}
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}

	//flush the data to ensure all data is written
	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}
	return csvData.Bytes(), nil
}
