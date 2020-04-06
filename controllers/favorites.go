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
	type Department struct {
		Title string        `json:"title"`
		Items []models.Item `json:"items"`
	}
	favDepts   := []string{
	    "Fruits & Vegetables",
    	"Snacks & Candy",
    	"Household Essentials",
    	"Beverages",
    	"Meat",
    	"Frozen",
    	"Eggs & Dairy",
    	"Pantry",
    	"Baby",
    	"Bread & Bakery",
    	"Groceries & Household Essentials",
	}

	// Get this user's favorites
	var favs []models.Favorite
	db.Where("user_id = ?", userId).Find(&favs)
	favMap := make(map[uint]bool)
	for _, fav := range favs { // fav map {itemId1: true, itemId2: true}
		favMap[fav.ItemId] = true
	}

	// Set up departments
	var favorites []Department
	rows, _  := db.Raw(`SELECT items.* from items, favorites 
						WHERE  items.id = favorites.item_id and favorites.user_id = ?
						ORDER  by items.id asc`, userId).Rows()
	defer rows.Close()
	deptMap   := make(map[string]int) // map deptName => index in favorites
	count     := 0
	numDepts  := len(favDepts)
	for rows.Next() {
		var item models.Item
		db.ScanRows(rows, &item)
		if _, ok := favMap[item.ID]; ok {
			item.Favorite = true
		} else {
			item.Favorite = false
		}
		
		// randIdx  := rand.Intn(len(favDepts))
		randDept := favDepts[count % numDepts]
		if _, ok := deptMap[randDept]; ok {
			// department exists in result, so append item
			deptIndex := deptMap[randDept]
			favorites[deptIndex].Items = append(favorites[deptIndex].Items, item)
		} else {
			// department does not exist, add to result + add item
			deptMap[randDept] = len(favorites)
			items := []models.Item{}
			items  = append(items, item)
			dept  := Department{Title: randDept, Items: items}
			favorites = append(favorites, dept)
		}
		count++
	}
	c.JSON(http.StatusOK, gin.H{"favorites": favorites})
}

// POST /favorites/:item_id
func InsertOrDeleteFavorites(c *gin.Context) {
	db         := c.MustGet("db").(*gorm.DB)
	action     := c.DefaultQuery("action", "insert")
	var favorite models.Favorite
	err := c.BindJSON(&favorite)
	if err != nil {
		log.Print(err)
	}
	if action == "delete" {
		// result := db.Where("item_id = ? and user_id = ?", favorite.ItemId, favorite.UserId).Delete(Favorite{})
		// result := db.Delete(Favorite{}, "item_id = ? and user_id = ?", favorite.ItemId, favorite.UserId)
		rows, err  := db.Raw(`DELETE from favorites
	                          WHERE  user_id = ? and item_id = ?`, favorite.UserId, favorite.ItemId).Rows()
		defer rows.Close()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "delete unsuccessfull"})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": ""})
		}
	} else {
		db.Create(&favorite)
		fail      := db.NewRecord(favorite) // check if insert succeeded
		if ! fail {
			c.JSON(http.StatusOK, gin.H{"error": ""})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "record exists"})
		}
	}
}

// GET /item/:id
func IsFavorite(c *gin.Context) {
	db     := c.MustGet("db").(*gorm.DB)
	id, _  := strconv.ParseInt(c.Param("id"), 10, 64)
	var item models.Item
	db.Where("id = ?", id).First(&item)
	c.JSON(http.StatusOK, gin.H{"item": item})
}
