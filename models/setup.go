// models/setup.go

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost)/DemoApp?charset=utf8&parseTime=True&loc=Local")
  	if err != nil {
    	panic("Failed to connect to database!")
  	}

	// db.AutoMigrate(&Item{})
	db.AutoMigrate(&Favorite{})
	db.AutoMigrate(&OrderItem{})
	db.AutoMigrate(&Order{})

  	return db
}