package controllers

import (
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/jenniekibiri/go-backend/initializers"
	"github.com/jenniekibiri/go-backend/models"
)

// func to create a restaurant
func CreateRestaurant(c *gin.Context) {
	// get the body of the request
	var body struct {
		RestaurantName string
		Address        string
		Photo          string
		Lat            float64
		Long           float64
		Rating         int
	}
	// check if empty or valide
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}

	// create a restaurant
	fmt.Println(body)
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
	initializers.DB.Create(&restaurant)

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
		AuthorName string
		Rating     int
		Text       string
	}
	// check if empty or valide
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}
	// get the restaurant id
	restaurantId := c.Param("rating")
	// get the restaurant
	var restaurant models.Restaurant
	initializers.DB.Preload("Ratings").Where("id = ?", restaurantId).First(&restaurant)
	// create a rating
	rating := models.Rating{
		AuthorName:   body.AuthorName,
		Rating:       body.Rating,
		Text:         body.Text,
		RestaurantID: restaurant.ID,
	}
	// save the rating
	initializers.DB.Create(&rating)
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
