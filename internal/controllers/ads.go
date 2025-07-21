package controllers

import (
	"Marketplace/internal/config"
	"Marketplace/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAd(c *gin.Context) {
	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		ImageURL    string  `json:"image_url"`
		Price       float64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	ad := models.Ad{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Price:       input.Price,
		UserID:      userID,
	}

	config.DB.Create(&ad)

	c.JSON(http.StatusOK, ad)
}

func GetAds(c *gin.Context) {
	var ads []models.Ad

	order := "created_at desc"
	if c.Query("sort") == "price_asc" {
		order = "price asc"
	} else if c.Query("sort") == "price_desc" {
		order = "price desc"
	}

	minPrice, _ := strconv.ParseFloat(c.Query("min_price"), 64)
	maxPrice, _ := strconv.ParseFloat(c.Query("max_price"), 64)

	query := config.DB.Order(order)
	if minPrice != 0 {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice != 0 {
		query = query.Where("price <= ?", maxPrice)
	}

	query.Find(&ads)

	c.JSON(http.StatusOK, ads)
}
