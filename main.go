package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jenniekibiri/go-backend/controllers"
	"github.com/jenniekibiri/go-backend/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDb()

}
func main() {
	r:=gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.POST("/restaurants", controllers.CreateRestaurant)
	r.GET("/restaurants", controllers.GetRestaurants)
	r.GET("/restaurants/:rating", controllers.FilterRestaurantsByRating)
	r.POST("/restaurant/:id/", controllers.AddRatingToRestaurant)
	r.GET("/restaurant/:id/", controllers.GetReviewsOfRestaurant)

	r.Run(":4000")
}