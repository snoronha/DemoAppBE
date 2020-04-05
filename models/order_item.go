package models

import (
	_ "github.com/jinzhu/gorm"
)

type OrderItem struct {
	ID                 uint    `json:"id" gorm:"primary_key"`
	OrderId            uint    `json:"orderId"`
	ItemId             uint    `json:"itemId"`
	Quantity           uint    `json:"quantity"`
}