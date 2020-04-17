package models

import (
	_ "github.com/jinzhu/gorm"
)

type Store struct {
	ID                  uint    `json:"id" gorm:"primary_key"`
	Lat                 float64 `json:"lat"`
	Lng                 float64 `json:"lng"`
	Viewport            string  `json:"viewport" gorm:"type:varchar(255)"`
	Types               string  `json:"types"`
	Vicinity            string  `json:"vicinity" gorm:"type:varchar(255);unique_index"`
}