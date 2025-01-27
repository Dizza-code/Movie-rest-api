package routes

import (
	"example.com/movies-api/controllers"
	"example.com/movies-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	movies := server.Group("/movies")
	// movies.POST("/", controllers.CreateMovie)
	// movies.PUT("/:id", controllers.UpdateMovie)
	// movies.DELETE("/:id", controllers.DeleteMovie)
	// movies.DELETE("/", controllers.DeleteAllMovies)
	movies.GET("/", controllers.ListAllMovies)
	movies.GET("/one/:name", controllers.FindMovieByName)
	movies.GET("/all/:name", controllers.FindAllMoviesByName)
	// movies.POST("/multiple", controllers.InsertMultipleMovies)
	movies.GET("/:id", controllers.GetMovieByID)
	movies.GET("/download-csv", controllers.DownloadMovieCSV) // Add this line

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("movies/", controllers.CreateMovie)
	authenticated.PUT("movies/:id", controllers.UpdateMovie)
	authenticated.DELETE("movies/:id", controllers.DeleteMovie)
	authenticated.DELETE("movies/", controllers.DeleteAllMovies)
	authenticated.POST("/movies/multiple", controllers.InsertMultipleMovies)

	server.POST("/signup", controllers.SignUp)
	server.POST("/login", controllers.LoginUser)
}
