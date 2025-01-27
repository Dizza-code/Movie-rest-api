package controllers

import (
	"net/http"

	"example.com/movies-api/models"
	"github.com/gin-gonic/gin"
)

func DownloadMovieCSV(c *gin.Context) {
	// Query the movies with user email
	movies, err := models.GetMoviesWithUserEmail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movies"})
		return
	}
	// Generate the CSV data
	csvData, err := models.GenerateCSV(movies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CSV"})
		return
	}
	// Set the response headers
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=movies.csv")

	// Send the CSV data as the response
	c.Data(http.StatusOK, "text/csv", csvData)
}
