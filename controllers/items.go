package controllers

import (
	"DemoAppBE/models"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/reiver/go-porterstemmer"
)

// GET /home carousels
func HomeItems(c *gin.Context) {
	type Carousel struct {
		Title string        `json:"title"`
		Items []models.Item `json:"items"`
	}
	carouselTitles := []string{
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
	limit    := 50
	randPage := 1 + rand.Intn(20)
	offset   := (randPage - 1) * limit
	var items []models.Item
	var carousels []Carousel
	db.Limit(limit).Offset(offset).Find(&items)

	// Get this user's favorites
	userId := 1  // this will be an input param
	var favs []models.Favorite
	db.Where("user_id = ?", userId).Find(&favs)
	favMap := make(map[uint]bool)
	for _, fav := range favs { // fav map {itemId1: true, itemId2: true}
		favMap[fav.ItemId] = true
	}
	count := 0;
	for i := 0; i < 8; i++ {
		carousel := new(Carousel)
		carousel.Title = carouselTitles[i]
		carouselItems := new([]models.Item)
		for j := 0; j < 6; j++ {
			if _, ok := favMap[items[count].ID]; ok {
				items[count].Favorite = true
			} else {
				items[count].Favorite = false
			}
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
	  
	// Get this user's favorites and merge in
	userId := 1  // this will be an input param
	var favs []models.Favorite
	db.Where("user_id = ?", userId).Find(&favs)
	favMap := make(map[uint]bool)
	for _, fav := range favs { // fav map {itemId1: true, itemId2: true}
		favMap[fav.ItemId] = true
	}
	for idx, item := range items {
		if _, ok := favMap[item.ID]; ok {
			items[idx].Favorite = true
		} else {
			items[idx].Favorite = false
		}
	}

  	c.JSON(http.StatusOK, gin.H{"items": items})
}

// GET /item/:id
func FindItem(c *gin.Context) {
	db     := c.MustGet("db").(*gorm.DB)
	id, _  := strconv.ParseInt(c.Param("id"), 10, 64)
	var item models.Item
	db.Where("id = ?", id).First(&item)
	c.JSON(http.StatusOK, gin.H{"item": item})
}

// GET /items/search?kwd=<keyword>
// Search for item
func SearchItems(c *gin.Context) {
	db         := c.MustGet("db").(*gorm.DB)
	kwd        := c.DefaultQuery("kwd", "apple")
	stemmed    := strings.ToLower(porterstemmer.StemString(kwd))
	searchTerm := "%" + stemmed + "%"
	var items []models.Item = make([]models.Item, 0)
	if len(stemmed) > 0 {
		db.Where("name_lc like ?", searchTerm).Find(&items)

		// Get this user's favorites and merge in
		userId := 1  // this will be an input param
		var favs []models.Favorite
		db.Where("user_id = ?", userId).Find(&favs)
		favMap := make(map[uint]bool)
		for _, fav := range favs { // fav map {itemId1: true, itemId2: true}
			favMap[fav.ItemId] = true
		}
		for idx, item := range items {
			if _, ok := favMap[item.ID]; ok {
				items[idx].Favorite = true
			} else {
				items[idx].Favorite = false
			}
		}
	}
	
	c.JSON(http.StatusOK, gin.H{"count": len(items), "items": items})
}