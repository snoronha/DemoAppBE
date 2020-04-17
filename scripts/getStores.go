package main

import (
	"DemoAppBE/models"
	"DemoAppBE/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type Location struct {
	Lat                *float64    `json:"lat"`
	Lng                *float64    `json:"lng"`
}
type Geom struct {
	Location           *Location   `json:"location"`
	Viewport           interface{} `json:"viewport"` 
}
// Use pointers for all nullable fields (else defaults to 0, 0.0, "" in Go)
type StoreObj struct {
	Geometry           *Geom       `json:"geometry"`
	Types              interface{} `json:"types"`
	Vicinity           *string     `json:"vicinity"`
}

type Results struct {
	Results            []StoreObj `json:"results"`
}

/*
	{
   "html_attributions" : [],
   "next_page_token" : "CqQCHAEAADTDKhFEkmDUX2YQp6V9JmHvvUm45S2Cinr6scQz_VW8qvBT66uI-Tywskq8uDcDoZsPjykz_p4Nwa-H2n4R4i_u-IvGEGFxvV1TRJjsuo9ou6fD8_ZdkBPHe_1Bf4DIyMde3nrvpHymB4PhkPfXqVxZYYMtf3tj1dFwEVvLQBjNjFSkZ-6D2HIakWnFcbksc1Tc6iJWMgzYY7vKJimTun8nilmoY5Osj0_BqxzaIRIDi_VYZ1OnURLauA-olTeIFazs6oHF5B0BncQIeT39TZzYCsuqj8kuXeiUNQTvPvj5Oz6IMemalELGw2P2AjNnlOd7-i8cbYX2NBvMjFvy08lkgllypOkd7g_5OkjZQD_JR39qWyngVVxUlOOYhFH5XBIQZcGqC9X4UKxwR2b1pUtr2RoUvpogCl_vCHPnEHa81XamRG8xHSI",
   "results" : [
      {
         "geometry" : {
            "location" : {
               "lat" : 37.4007514,
               "lng" : -122.1098392
            },
            "viewport" : {
               "northeast" : {
                  "lat" : 37.40226552989272,
                  "lng" : -122.1080340701073
               },
               "southwest" : {
                  "lat" : 37.39956587010727,
                  "lng" : -122.1107337298928
               }
            }
         },
         "icon" : "https://maps.gstatic.com/mapfiles/place_api/icons/shopping-71.png",
         "id" : "5aeb514a2b83e418b285fa050e95123037821493",
         "name" : "Walmart",
         "opening_hours" : {
            "open_now" : true
         },
         "photos" : [
            {
               "height" : 4032,
               "html_attributions" : [
                  "\u003ca href=\"https://maps.google.com/maps/contrib/117419991471593763887\"\u003ebeem parthiban\u003c/a\u003e"
               ],
               "photo_reference" : "CmRaAAAAW-B5kyuGZLjC3vR3yXXWUFW1hvwKZ8YtVtDIkcrleBAHsVSHLXH0dmoM35AVPcpkXkxO0Bg658EAj5NjwnjMqWWBU5S96Q4X_FCeNgOQkhl5Uqf3xPjg-rFwdNzRjHSYEhAx_OinK97HmypJTbBpOiwzGhTRHE6_UdLSCsrE67-_YUc9dz9yTw",
               "width" : 3024
            }
         ],
         "place_id" : "ChIJryH6Apmwj4ARx08eJ5hAaGQ",
         "plus_code" : {
            "compound_code" : "CV2R+83 Mountain View, California",
            "global_code" : "849VCV2R+83"
         },
         "price_level" : 1,
         "rating" : 3.7,
         "reference" : "ChIJryH6Apmwj4ARx08eJ5hAaGQ",
         "scope" : "GOOGLE",
         "types" : [
            "department_store",
            "supermarket",
            "grocery_or_supermarket",
            "food",
            "point_of_interest",
            "store",
            "establishment"
         ],
         "user_ratings_total" : 3928,
         "vicinity" : "600 Showers Dr, Mountain View"
	  },
	  ...
	]
	}

*/

func main() {
	db       := util.GetDB()
	db.AutoMigrate(&models.Store{})

	/*
	fullPath := os.Getenv("GOPATH") + "/src/DemoAppBE"
    absFile, err := filepath.Abs(fullPath + "/scripts/output")
    if err != nil {
        panic("Failed to get absolute path")
	}
	fmt.Println("FILE = " + absFile)
    body, fileErr := ioutil.ReadFile(absFile)
    if fileErr != nil {
        panic("Failed to read output")
	}
	*/
	
	client   := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}
	endpoint  := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
	APIKey    := "AIzaSyDE49x8p7SZG4p2nNJDUnk5uYMkJuPEBnQ"
	locations := []string{
		"37.282057,-122.040169",   // San Jose
		"38.5616505,-121.5829967", // Sacramento
		"34.0201613,-118.6919201", // Los Angeles
		"32.8242404,-117.3891666", // San Diego
		"31.8110563,-106.5646024", // El Paso
		"29.481137,-98.7945941",   // San Antonio
		"30.3076863,-97.8934865",  // Austin
		"32.8203525,-97.0117306",  // Dallas
		"29.8168824,-95.6814809",  // Houston
		"36.3674019,-94.308865",   // Bentonville
		"30.3446913,-81.9632962",  // Jacksonville
		"28.4810971,-81.5089238",  // Orlando
		"27.9944116,-82.5943658",  // Tampa
		"25.7823907,-80.2994988",  // Miami
	}

	storeMap := make(map[string]models.Store)
	for _, location := range locations {
		fmt.Printf("Doing %s ...\n", location)
		url  := endpoint + "?location=" + location + "&radius=50000&keyword=walmart&key=" + APIKey
		body := util.GetUrltext(url, client)
		var stores Results
		jsonErr  := json.Unmarshal(body, &stores)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		for _, objStore := range stores.Results {
			store := flattenStore(objStore)
			if _, ok := storeMap[store.Vicinity]; !ok {
				// Store does not exist, insert into storeMap
				storeMap[store.Vicinity] = store
			} else {
				fmt.Printf("Seen before: %s\n", store.Vicinity)
			}
		}
	}

	// save itemMap to db
	fmt.Printf("Saving to Store ...\n")
	saveStoresToDB(db, storeMap)
}

func saveStoresToDB(db *gorm.DB, storeMap map[string]models.Store) {
	for _, store := range storeMap {
		db.Create(&store)
	}
}

func flattenStore(store StoreObj) models.Store {
	viewport, _ := json.Marshal(store.Geometry.Viewport)
	types, _    := json.Marshal(store.Types)
	return models.Store{
		Lat:                *store.Geometry.Location.Lat,
		Lng:                *store.Geometry.Location.Lng,
		Viewport:           string(viewport),
		Types:              string(types),
		Vicinity:           *store.Vicinity,
	
	}
}
