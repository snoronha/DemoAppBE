package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/reiver/go-porterstemmer"
)

func GetDB() *gorm.DB {
    file, _ := ioutil.ReadFile("../conf.json")
    type Conf struct{
        DBUser     string
        DBPassword string
        DBHost     string
        DBName     string
    }

    conf := Conf{}
    _ = json.Unmarshal([]byte(file), &conf)

    dbStr   := fmt.Sprintf(
        "%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", 
        conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBName,
    )

	db, err := gorm.Open("mysql", dbStr)
	if err != nil {
	  panic("Failed to connect: " + err.Error())
    }
    return db
}

func GetUrltext(url string, client http.Client) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}

func StemSentence(str string) string {
	reg, err    := regexp.Compile("[^a-zA-Z]+")
	strs        := strings.Fields(str)
	if err != nil {
        log.Fatal(err)
    }
	stemmedStrs := []string{}
	for _, word := range strs {
		stemmedWord := reg.ReplaceAllString(word, "")
		stemmedWord  = porterstemmer.StemString(strings.ToLower(stemmedWord))
		stemmedStrs  = append(stemmedStrs, stemmedWord)
	}
	return strings.Join(stemmedStrs[:], " ")
}