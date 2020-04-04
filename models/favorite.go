package models

import (
	_ "github.com/jinzhu/gorm"
)

type Favorite struct {
	ID                 uint    `json:"id" gorm:"primary_key"`
	ItemId             uint    `json:"itemId"`
    UserId             uint    `json:"userId"`
}