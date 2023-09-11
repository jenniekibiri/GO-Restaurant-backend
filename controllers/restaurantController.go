package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jenniekibiri/go-backend/initializers"
	"github.com/jenniekibiri/go-backend/models"
)

// @Summary Create a new restaurant
// @Description Create a new restaurant entry
// @Accept  json
// @Produce  json
// @Param restaurantName body string true "Restaurant Name"
// @Param address body string true "Address"
// @Param photo body string true "Photo"
// @Param lat body float64 true "Latitude"
// @Param long body float64 true "Longitude"
// @Param rating body int true "Rating"
// @Success 200 {object} string "Restaurant created successfully"
// @Failure 400 {object} string "Fields are empty"
// @Failure 500 {object} string "Failed to save restaurant"
// @Router /restaurants [post]
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

// @Summary Get all restaurants
// @Description  Get all restaurants
// @Accept  json
// @Produce  json
// @Router /restaurants [get]
func GetRestaurants(c *gin.Context) {
	// get all restaurants
	var restaurants []models.Restaurant
	initializers.DB.Preload("Ratings").Find(&restaurants)
	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

// @Summary Filter restaurants by rating
// @Description  Filter restaurants by rating
// @Accept  json
// @Produce  json
// @Param rating path string true "Rating"
// @Router /restaurants/{rating} [get]
func FilterRestaurantsByRating(c *gin.Context) {
	// get the rating
	rating := c.Param("rating")
	// get all restaurants with that rating
	var restaurants []models.Restaurant
	initializers.DB.Preload("Ratings").Where("rating = ?", rating).Find(&restaurants)
	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

// @Summary Add rating to a restaurant
// @Description  Add rating to a restaurant
// @Accept  json
// @Produce  json
// @Param id path string true "Restaurant ID"
// @Param author_name body string true "Author Name"
// @Param rating body int true "Rating"
// @Param text body string true "Text"
// @Router /restaurant/{id} [post]
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

// @Summary Get reviews of a restaurant
// @Description  Get reviews of a restaurant
// @Accept  json
// @Produce  json
// @Param id path string true "Restaurant ID"
// @Router /restaurant/{id} [get]
func GetReviewsOfRestaurant(c *gin.Context) {
	// get the restaurant id
	restaurantId := c.Param("id")
	// get the restaurant
	var restaurant models.Restaurant
	initializers.DB.Preload("Ratings").Where("id = ?", restaurantId).First(&restaurant)
	// return the restaurant
	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}