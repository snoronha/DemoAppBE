package controllers

import (
	"net/http"
  	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
  	"DemoAppBE/models"
)

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