package controllers

import (
	"net/http"
	"strings"
  	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/reiver/go-porterstemmer"
	"strconv"
  	"DemoAppBE/models"
)

// GET /home carousels
func HomeItems(c *gin.Context) {
	type Carousel struct {
		Title string        `json:"title"`
		Items []models.Item `json:"items"`
	}
	carouselTitles := [8]string{
    	"Featured Items",
    	"Reorder Your Essentials",
    	"Healthy Snacking",
    	"Easy Cleanup",
    	"Recommended for You",
    	"Fresh Fruit",
    	"Beef",
    	"Nuts & Dried Fruit",
	}
	db       := c.MustGet("db").(*gorm.DB)
	var items []models.Item
	var carousels []Carousel
	db.Limit(50).Find(&items)
	count := 0;
	for i := 0; i < 8; i++ {
		carousel := new(Carousel)
		carousel.Title = carouselTitles[i]
		carouselItems := new([]models.Item)
		for j := 0; j < 6; j++ {
			*carouselItems = append(*carouselItems, items[count])
			count++
		}
		carousel.Items = *carouselItems
		carousels = append(carousels, *carousel)
	}

	c.JSON(http.StatusOK, gin.H{"carousels": carousels})
}

// GET /items
// Get all items
func FindItems(c *gin.Context) {
  	db       := c.MustGet("db").(*gorm.DB)
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "60"), 10, 64)
	page, _  := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	offset   := (page - 1) * limit
  	var items []models.Item
  	db.Limit(limit).Offset(offset).Order("id asc").Find(&items)

  	c.JSON(http.StatusOK, gin.H{"items": items})
}

// GET /items/search?kwd=<keyword>
// Search for item
func SearchItems(c *gin.Context) {
	db         := c.MustGet("db").(*gorm.DB)
	kwd        := c.DefaultQuery("kwd", "apple")
	stemmed    := strings.ToLower(porterstemmer.StemString(kwd))
	searchTerm := "%" + stemmed + "%"
	var items []models.Item
	db.Where("name_lc like ?", searchTerm).Find(&items)

	c.JSON(http.StatusOK, gin.H{"count": len(items), "items": items})
}