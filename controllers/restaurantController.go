package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jenniekibiri/go-backend/initializers"
	"github.com/jenniekibiri/go-backend/models"
)

// func to create a restaurant
func CreateRestaurant(c *gin.Context) {
	// get the body of the request
	var body struct {
		RestaurantName string  `binding:"required"`
		Address        string  `binding:"required"`
		Photo          string  `binding:"required"`
		Lat            float64 `binding:"required"`
		Long           float64 `binding:"required"`
		Rating         int     `binding:"required"`
	}
	// check if empty or valide
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}

	// create a restaurant

	restaurant := models.Restaurant{
		RestaurantName: body.RestaurantName,
		Address:        body.Address,
		Photo:          body.Photo,
		Lat:            body.Lat,
		Long:           body.Long,
		Rating:         body.Rating,
		Ratings:        []models.Rating{},
	}

	// save the restaurant
	result :=initializers.DB.Create(&restaurant)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save rating"})
		return
	}

	// return the restaurant
	c.JSON(http.StatusOK, gin.H{"data": restaurant})

}

// function to get all restaurants
func GetRestaurants(c *gin.Context) {
	// get all restaurants
	var restaurants []models.Restaurant
	initializers.DB.Preload("Ratings").Find(&restaurants)
	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

// function to filter rastaurants by rating
func FilterRestaurantsByRating(c *gin.Context) {
	// get the rating
	rating := c.Param("rating")
	// get all restaurants with that rating
	var restaurants []models.Restaurant
	initializers.DB.Preload("Ratings").Where("rating = ?", rating).Find(&restaurants)
	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

// function to add a rating to a restaurant
func AddRatingToRestaurant(c *gin.Context) {
	var body struct {
		AuthorName string `json:"author_name" binding:"required"`
		Rating     int    `json:"rating" binding:"required"`
		Text       string `json:"text" binding:"required"`
	}
	// check if empty or valide
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}
	// get the restaurant id
	restaurantId := c.Param("id")

	// get the restaurant
	var restaurant models.Restaurant
	initializers.DB.Preload("Ratings").Where("id = ?", restaurantId).First(&restaurant)
	// Check if the restaurant exists
	if restaurant.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	// create a rating
	rating := models.Rating{
		AuthorName:   body.AuthorName,
		Rating:       body.Rating,
		Text:         body.Text,
		RestaurantID: restaurant.ID,
	}
	// save the rating
	result := initializers.DB.Create(&rating)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save rating"})
		return
	}
	// return the restaurant
	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}

// function to get reviews of a restaurant
func GetReviewsOfRestaurant(c *gin.Context) {
	// get the restaurant id
	restaurantId := c.Param("id")
	// get the restaurant
	var restaurant models.Restaurant
	initializers.DB.Preload("Ratings").Where("id = ?", restaurantId).First(&restaurant)
	// return the restaurant
	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}
