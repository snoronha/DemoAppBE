package models

import (
	_ "github.com/jinzhu/gorm"
)

type Item struct {
    ID                 uint    `json:"id" gorm:"primary_key"`
    UsItemId           string  `json:"USItemId" sql:"type:varchar(16)"`
    OfferId            string  `json:"offerId" gorm:"type:varchar(64)"`
    Name               string  `json:"name"`
    Thumbnail          string  `json:"thumbnail"`
    WeightIncrement    float64 `json:"weightIncrement"`
    AverageWeight      float64 `json:"averageWeight"`
    MaxAllowed         uint    `json:"maxAllowed"`
    ProductUrl         string  `json:"productUrl"`
    IsSnapEligible     uint    `json:"isSnapEligible"`
    Type               string  `json:"type"`
    Rating             float64 `json:"rating"`
    ReviewsCount       uint    `json:"reviewsCount"`
    IsOutOfStock       string  `json:"isOutOfStock"`
    List               float64 `json:"list"`
    PreviousPrice      float64 `json:"previousPrice"`
    PriceUnitOfMeasure string  `json:"priceUnitOfMeasure"`
    SalesUnitOfMeasure string  `json:"salesUnitOfMeasure"`
    SalesQuantity      uint    `json:"salesQuantity"`
    DisplayCondition   string  `json:"displayCondition"`
    DisplayPrice       float64 `json:"displayPrice"`
    DisplayUnitPrice   string  `json:"displayUnitPrice"`
    IsClearance        string  `json:"isClearance"`
    IsRollback         string  `json:"isRollback"`
    Unit               string  `json:"unit"`
}