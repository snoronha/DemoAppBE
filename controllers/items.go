package controllers

import (
	"net/http"
  	"github.com/gin-gonic/gin"
  	"github.com/jinzhu/gorm"
  	"DemoAppBE/models"
)

// GET /items
// Get all items
func FindItems(c *gin.Context) {
  	db := c.MustGet("db").(*gorm.DB)

  	var items []models.Item
  	db.Find(&items)

  	c.JSON(http.StatusOK, gin.H{"data": items})
}