package main

import (
    "github.com/gin-gonic/gin"
    "DemoAppBE/models"
    "DemoAppBE/controllers"
)

func main() {
    router := gin.Default()

    db := models.SetupModels() // new

    // Provide db variable to controllers
    router.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    router.GET("/home", controllers.HomeItems)
    router.GET("/items", controllers.FindItems) // new
    router.GET("/items/search", controllers.SearchItems)

    router.Run()
}