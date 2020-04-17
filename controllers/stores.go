package controllers

import (
	"DemoAppBE/models"
	"DemoAppBE/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /stores/:lat/:lng
func GetStores(c *gin.Context) {
	db       := c.MustGet("db").(*gorm.DB)
	lat1, _  := strconv.ParseFloat(c.Param("lat"), 64)
	lng1, _  := strconv.ParseFloat(c.Param("lng"), 64)

	var MaxRadius float64 = 50000.0 // 50km
	var stores []models.Store
	db.Find(&stores)
	nearbyStores := []models.Store{}
	for _, store := range stores {
		lat2 := store.Lat
		lng2 := store.Lng
		if util.Distance(lat1, lng1, lat2, lng2) <= MaxRadius {
			nearbyStores = append(nearbyStores, store)
		}
	}
	c.JSON(http.StatusOK, gin.H{"item_detail": nearbyStores})
}
