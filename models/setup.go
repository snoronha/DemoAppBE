// models/setup.go

package models

import (
	"DemoAppBE/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db := util.GetDB()

	// db.AutoMigrate(&Item{})
	db.AutoMigrate(&Favorite{})
	db.AutoMigrate(&OrderItem{})
	db.AutoMigrate(&Order{})

  	return db
}