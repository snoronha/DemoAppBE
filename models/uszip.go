package models

import (
	_ "github.com/jinzhu/gorm"
)

type USZip struct {
	ID                  uint    `json:"id" gorm:"primary_key"`
	Zip                 string  `json:"zip" gorm:"type:varchar(16);unique_index;not null"`
	Lat                 float64 `json:"lat"`
	Lng                 float64 `json:"lng"`
	City                string  `json:"city" gorm:"type:varchar(32);index:city_idx"`
	StateID             string  `json:"stateId" gorm:"type:varchar(8)"`
	StateName           string  `json:"stateName" gorm:"type:varchar(32)"`
	Population          uint    `json:"population"`
	CountyName          string  `json:"countyName" gorm:"type:varchar(32)"`
	TimeZone            string  `json:"timeZone" gorm:"type:varchar(32)"`
}