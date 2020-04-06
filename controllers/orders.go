package controllers

import (
	"DemoAppBE/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /order/:order_id/user/:user_id
func ReadOrder(c *gin.Context) {
	db         := c.MustGet("db").(*gorm.DB)
	orderId, _ := strconv.ParseInt(c.Param("order_id"), 10, 64)
	userId, _  := strconv.ParseInt(c.Param("user_id"), 10, 64)

	// Get this user's favorites
	var favs []models.Favorite
	db.Where("user_id = ?", userId).Find(&favs)
	favMap := make(map[uint]bool)
	for _, fav := range favs { // fav map {itemId1: true, itemId2: true}
		favMap[fav.ItemId] = true
	}
	// Get cart by order_id
	rows, _  := db.Raw(`SELECT items.*, order_items.quantity, order_items.order_id
	                    FROM   items, order_items 
						WHERE  items.id = order_items.item_id and 
							   order_items.order_id = ?`, orderId).Rows()				   
	defer rows.Close()
	orderItems := []models.Item{}
	for rows.Next() {
		var item models.Item
		db.ScanRows(rows, &item)
		if _, ok := favMap[item.ID]; ok {
			item.Favorite = true
		} else {
			item.Favorite = false
		}
		orderItems = append(orderItems, item)
	}
	c.JSON(http.StatusOK, gin.H{"order": orderItems})
}

// POST /order_item/:order_id
func UpsertOrderItem(c *gin.Context) {
	db         := c.MustGet("db").(*gorm.DB)
	orderId, _ := strconv.ParseInt(c.Param("order_id"), 10, 64)
	var order_item models.OrderItem
	err := c.BindJSON(&order_item)
	if err != nil {
		log.Print(err)
	}
	order      := models.Order{Status: "open"}
	db.Where(models.Order{ID: uint(orderId)}).FirstOrCreate(&order)
	if order_item.Quantity <= 0 {
		rows, _  := db.Raw(`DELETE from order_items
	                        WHERE  order_id = ? and item_id = ?`, order_item.OrderId, order_item.ItemId).Rows()
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{"order_item": order_item})
	} else {
		db.Where(models.OrderItem{ItemId: uint(order_item.ItemId)}).Assign(models.OrderItem{Quantity: order_item.Quantity}).FirstOrCreate(&order_item)
		c.JSON(http.StatusOK, gin.H{"order_item": order_item})
	}
}

/*
// GET /item/:id
func IsFavorite(c *gin.Context) {
	db     := c.MustGet("db").(*gorm.DB)
	id, _  := strconv.ParseInt(c.Param("id"), 10, 64)
	var item models.Item
	db.Where("id = ?", id).First(&item)
	c.JSON(http.StatusOK, gin.H{"item": item})
}
*/
