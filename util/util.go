package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/reiver/go-porterstemmer"
)

func GetDB() *gorm.DB {
    fullPath := os.Getenv("GOPATH") + "/src/DemoAppBE"
    absFile, err := filepath.Abs(fullPath + "/conf.json")
    if err != nil {
        panic("Failed to get absolute path")
    }
    file, fileErr := ioutil.ReadFile(absFile)
    if fileErr != nil {
        panic("Failed to read conf.json")
    }
    type Conf struct{
        DBUser     string
        DBPassword string
        DBHost     string
        DBName     string
    }

    conf := Conf{}
    err   = json.Unmarshal([]byte(file), &conf)
    if err != nil {
        panic("Failed to read unmarshal json")
    }

    dbStr   := fmt.Sprintf(
        "%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", 
        conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBName,
    )
    fmt.Println("DBSTR = " + dbStr) 
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