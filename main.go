package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jenniekibiri/go-backend/controllers"
	"github.com/jenniekibiri/go-backend/initializers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/jenniekibiri/go-backend/docs"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDb()

}


// add swagger documentation 
// @title Restaurant API
// @description This is a sample restaurant API with Swagger documentation.
// @version 1
// @host localhost:4000
// @BasePath /
// @schemes http
func main() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/restaurants", controllers.CreateRestaurant)
	r.GET("/restaurants", controllers.GetRestaurants)
	r.GET("/restaurants/:rating", controllers.FilterRestaurantsByRating)
	r.POST("/restaurant/:id/", controllers.AddRatingToRestaurant)
	r.GET("/restaurant/:id/", controllers.GetReviewsOfRestaurant)

	r.Run(":4000")
}
