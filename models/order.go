package models

import (
	_ "github.com/jinzhu/gorm"
)

type Order struct {
	ID                 uint    `json:"id" gorm:"primary_key"`
	Status             string  `json:"status"`
	Total              float64 `json:"total"`
}