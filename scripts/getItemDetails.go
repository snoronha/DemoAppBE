package main

import (
	"DemoAppBE/models"
	"DemoAppBE/util"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type PrimaryDepartment struct {
	Id                  *string     `json:"id"`
	Name                *string     `json:"name"`
}
type PrimaryAisle struct {
	Id                  *string     `json:"id"`
	Name                *string     `json:"name"`
}
type PrimaryShelf struct {
	Id                  *string     `json:"id"`
	Name                *string     `json:"name"`
}
type Image struct {
	Thumbnail           *string     `json:"thumbnail"`      //thumbnail
	Large               *string     `json:"large"`   
}
type Price struct {
	List                *float64    `json:"list"`
    PriceUnitOfMeasure  *string     `json:"priceUnitOfMeasure"`
    SalesUnitOfMeasure  *string     `json:"salesUnitOfMeasure"`
	SalesQuantity       *uint       `json:"salesQuantity"`
	IsRollback          *string     `json:"isRollback"`
	IsClearance         *string     `json:"isClearance"`
	Unit                *float64    `json:"unit"`
	DisplayPrice        *float64    `json:"displayPrice"`
	DisplayUnitPrice    *string     `json:"displayUnitPrice"`
}
type Store struct {
	Price               *Price      `json:"price"`
	IsInStock           *bool       `json:"isInStock"`
}
type Detailed struct {
	ProductCode         *string     `json:"productCode"`
	Brand               *string     `json:"brand"`
	ProductType         *string     `json:"productType"`
	ShortDescription    *string     `json:"shortDescription"`
	Description         *string     `json:"description"`
	ModelNum            *string     `json:"modelNum"`
	AssembledInCountryOfOrigin   *string `json:"assembledInCountryOfOrigin"`
	OriginOfComponents  *string     `json:"originOfComponents"`
	Ingredients         *string     `json:"ingredients"`
	AgeRestricted       *bool       `json:"ageRestricted"`
	StorageType         *string     `json:"storageType"`
	Weight              *string     `json:"weight"`
	Rating              *float64    `json:"rating"`
	ReviewsCount        *uint       `json:"reviewsCount"`
}
type Basic struct {
	Name                *string     `json:"name"`
	MaxAllowed          *uint       `json:"maxAllowed"`
	TaxCode             *string     `json:"taxCode"`
	IsOutOfStock        *string     `json:"isOutOfStock"`
	Image               *Image      `json:"image"`
	IsAlcoholic         *bool       `json:"isAlcoholic"`
	IsSnapEligible      *uint       `json:"isSnapEligible"`
	PrimaryShelf        *PrimaryShelf      `json:"primaryShelf"`
	PrimaryAisle        *PrimaryAisle      `json:"primaryAisle"`
	PrimaryDepartment   *PrimaryDepartment `json:"primaryDepartment"`
	SalesUnit           *string     `json:"salesUnit"`
	ProductUrl          *string     `json:"productUrl"`
	Type                *string     `json:"type"`
}

type ItemDetailObj struct {
	Sku                 *string     `json:"sku"`
    UsItemId            *string     `json:"UsItemId"`
	OfferId             *string     `json:"offerId"`
	Upc                 *string     `json:"upc"`
	Rank                *uint       `json:"rank"`
	Basic               *Basic      `json:"basic"`
	Detailed            *Detailed   `json:"detailed"`
	NutritionFacts      interface{} `json:"nutritionFacts"`
	Store               *Store      `json:"store"`
}

type ItemDetail struct {
	Sku                 *string     `json:"sku" sql:"type:varchar(32)"`
    UsItemId            *string     `json:"USItemId" sql:"type:varchar(16)"`
	OfferId             *string     `json:"offerId" gorm:"type:varchar(64)"`
	Upc                 *string     `json:"upc" gorm:"type:varchar(32)"`
	Rank                *uint       `json:"rank"`
	Name                *string     `json:"name"`
	MaxAllowed          *uint       `json:"maxAllowed"`
	TaxCode             *string     `json:"taxCode"`
	IsOutOfStock        *string     `json:"isOutOfStock"`
	Thumbnail           *string     `json:"thumbnail"`      // thumbnail
	Large               *string     `json:"large"`          // large image
	IsAlcoholic         *bool       `json:"isAlcoholic"`
	IsSnapEligible      *uint       `json:"isSnapEligible"`
	PrimaryShelf        *string     `json:"primaryShelf"`
	PrimaryAisle        *string     `json:"primaryAisle"`
	PrimaryDepartment   *string     `json:"primaryDepartment"`
	SalesUnit           *string     `json:"salesUnit"`
	ProductUrl          *string     `json:"productUrl"`
	Type                *string     `json:"type"`
	ProductCode         *string     `json:"productCode"`
	Brand               *string     `json:"brand"`
	ProductType         *string     `json:"productType"`
	ShortDescription    *string     `json:"shortDescription"`
	Description         *string     `json:"description"`
	ModelNum            *string     `json:"modelNum"`
	AssembledInCountryOfOrigin  *string `json:"assembledInCountryOfOrigin"`
	OriginOfComponents  *string     `json:"originOfComponents"`
	Ingredients         *string     `json:"ingredients"`
	AgeRestricted       *bool       `json:"ageRestricted"`
	StorageType         *string     `json:"storageType"`
	Weight              *string     `json:"weight"`
	Rating              *float64    `json:"rating"`
	ReviewsCount        *uint       `json:"reviewsCount"`
	NutritionFacts      string      `json:"nutritionFacts"`
	List                *float64    `json:"list"`
    PriceUnitOfMeasure  *string     `json:"priceUnitOfMeasure"`
    SalesUnitOfMeasure  *string     `json:"salesUnitOfMeasure"`
	SalesQuantity       *uint       `json:"salesQuantity"`
	IsRollback          *string     `json:"isRollback"`
	IsClearance         *string     `json:"isClearance"`
	Unit                *float64    `json:"unit"`
	DisplayPrice        *float64    `json:"displayPrice"`
	DisplayUnitPrice    *string     `json:"displayUnitPrice"`
	IsInStock           *bool       `json:"isInStock"`
}

// Use pointers for all nullable fields (else defaults to 0, 0.0, "" in Go)
func main() {
	db      := util.GetDB()
	db.AutoMigrate(&models.ItemDetail{})

	itemMap := make(map[string]bool)
	var itemIds []string
	db.Model(&models.Item{}).Pluck("us_item_id", &itemIds)
	fmt.Printf("Item count = %d\n", len(itemIds))
	for _, itemId := range itemIds {
		if _, ok := itemMap[itemId]; !ok {
			itemMap[itemId] = true
		}
	}

	client  := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
	}

	count := 0
	for itemId, _ := range itemMap {
		if count < len(itemIds) {
			if count % 10 == 0 {
				fmt.Printf("  Count = %d\n", count)
			}
			var itemDet models.ItemDetail
			db.Where("us_item_id = ?", itemId).First(&itemDet)
			if itemDet.ID > 0 {
				count++
				fmt.Printf("    Duplicate itemId: %s\n", itemId)
				continue
			}

			url  := "https://grocery.walmart.com/v3/api/products/" + itemId + "?itemFields=all&storeId=2119"
			body := util.GetUrltext(url, client)
			time.Sleep(time.Duration(500) * time.Millisecond)
			var itemDetailObj ItemDetailObj
			err := json.Unmarshal(body, &itemDetailObj)
			_ = err
			itemDetail := flattenItemDetail(itemDetailObj)
			db.Create(&itemDetail)
		}
		count++
	}
}

func flattenItemDetail(item ItemDetailObj) ItemDetail {
	nutritionFacts, _  := json.Marshal(item.NutritionFacts)
	if item.Detailed.Description != nil {
		descr := *item.Detailed.Description
		if len(descr) >= 2047 {
			*item.Detailed.Description = descr[:2047]
		}
	}
	if item.Detailed.ShortDescription != nil {
		shortDescr  := *item.Detailed.ShortDescription
		if len(shortDescr) >= 2047 {
			*item.Detailed.ShortDescription = shortDescr[:2047]
		}
	}
	if item.Detailed.Ingredients != nil {
		ingredients := *item.Detailed.Ingredients
		if len(ingredients) >= 511 {
			*item.Detailed.Ingredients = ingredients[:511]
		}
	}
	primaryShelf := ""
	if item.Basic.PrimaryShelf != nil {
		primaryShelf = *item.Basic.PrimaryShelf.Name
	}
	primaryAisle := ""
	if item.Basic.PrimaryAisle != nil {
		primaryAisle = *item.Basic.PrimaryAisle.Name
	}
	primaryDepartment := ""
	if item.Basic.PrimaryDepartment != nil {
		primaryDepartment = *item.Basic.PrimaryDepartment.Name
	}
	thumbnail := ""
	large     := ""
	if item.Basic.Image != nil {
		thumbnail = *item.Basic.Image.Thumbnail
		large     = *item.Basic.Image.Large
	}
	return ItemDetail{
		Sku:                item.Sku,
		UsItemId:           item.UsItemId,
		OfferId:            item.OfferId,
		Upc:                item.Upc,
		Rank:               item.Rank,
		Name:               item.Basic.Name,
		MaxAllowed:         item.Basic.MaxAllowed,
		TaxCode:            item.Basic.TaxCode,
		IsOutOfStock:       item.Basic.IsOutOfStock,
		Thumbnail:          &thumbnail,
		Large:              &large,
		IsAlcoholic:        item.Basic.IsAlcoholic,
		IsSnapEligible:     item.Basic.IsSnapEligible,
		PrimaryShelf:       &primaryShelf,
		PrimaryAisle:       &primaryAisle,
		PrimaryDepartment:  &primaryDepartment,
		SalesUnit:          item.Basic.SalesUnit,
		ProductUrl:         item.Basic.ProductUrl,
		Type:               item.Basic.Type,
		ProductCode:        item.Detailed.ProductCode,
		Brand:              item.Detailed.Brand,
		ProductType:        item.Detailed.ProductType,
		ShortDescription:   item.Detailed.ShortDescription,
		Description:        item.Detailed.Description,
		ModelNum:           item.Detailed.ModelNum,
		AssembledInCountryOfOrigin:  item.Detailed.AssembledInCountryOfOrigin,
		OriginOfComponents: item.Detailed.OriginOfComponents,
		Ingredients:        item.Detailed.Ingredients,
		AgeRestricted:      item.Detailed.AgeRestricted,
		StorageType:        item.Detailed.StorageType,
		Weight:             item.Detailed.Weight,
		Rating:             item.Detailed.Rating,
		ReviewsCount:       item.Detailed.ReviewsCount,
		NutritionFacts:     string(nutritionFacts),
		List:               item.Store.Price.List,
    	PriceUnitOfMeasure: item.Store.Price.PriceUnitOfMeasure,
    	SalesUnitOfMeasure: item.Store.Price.SalesUnitOfMeasure,
		SalesQuantity:      item.Store.Price.SalesQuantity,
		IsRollback:         item.Store.Price.IsRollback,
		IsClearance:        item.Store.Price.IsClearance,
		Unit:               item.Store.Price.Unit,
		DisplayPrice:       item.Store.Price.DisplayPrice,
		DisplayUnitPrice:   item.Store.Price.DisplayUnitPrice,
		IsInStock:          item.Store.IsInStock,
	}
}
