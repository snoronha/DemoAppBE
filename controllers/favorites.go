package controllers

import (
	"DemoAppBE/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /favorites
func ReadFavorites(c *gin.Context) {
	db         := c.MustGet("db").(*gorm.DB)
	userId, _  := strconv.ParseInt(c.Param("user_id"), 10, 64)
	
	var favorites []models.Item
	rows, _  := db.Raw(`SELECT items.* from items, favorites 
	                    WHERE  items.id = favorites.item_id and favorites.user_id = ?`, userId).Rows()
	defer rows.Close()
	for rows.Next() {
		var item models.Item
  		db.ScanRows(rows, &item)
  		favorites = append(favorites, item)
	}

	c.JSON(http.StatusOK, gin.H{"favorites": favorites})
}

// POST /favorites
func InsertFavorites(c *gin.Context) {
	db  := c.MustGet("db").(*gorm.DB)
	var favorite models.Favorite
	err := c.BindJSON(&favorite)
	if err != nil {
		log.Print(err)
	}
	db.Create(&favorite)
	fail      := db.NewRecord(favorite) // check if insert succeeded
	if ! fail {
		c.JSON(http.StatusOK, gin.H{"error": ""})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "record exists"})
	}
}

// GET /items
// Get all items
/*
func SaveFavorite(c *gin.Context) {
  	db       := c.MustGet("db").(*gorm.DB)
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "60"), 10, 64)
	page, _  := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	offset   := (page - 1) * limit
  	var items []models.Item
  	db.Limit(limit).Offset(offset).Order("id asc").Find(&items)

  	c.JSON(http.StatusOK, gin.H{"items": items})
}

// GET /item/:id
func IsFavorite(c *gin.Context) {
	db     := c.MustGet("db").(*gorm.DB)
	id, _  := strconv.ParseInt(c.Param("id"), 10, 64)
	var item models.Item
	db.Where("id = ?", id).First(&item)
	c.JSON(http.StatusOK, gin.H{"item": item})
}
*/
