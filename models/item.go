package models

import (
	_ "github.com/jinzhu/gorm"
)

type Item struct {
  ID     uint      `json:"id" gorm:"primary_key"`
  USItemId  string `json:"USItemId"`
  OfferId   string `json:"offerId"`
  Name      string `json:"name"`
  Thumbnail string `json:"thumbnail"`
}