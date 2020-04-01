package main

import (
    "github.com/gin-gonic/gin"
    "DemoAppBE/models"
    "DemoAppBE/controllers"
)

func main() {
    r := gin.Default()

    db := models.SetupModels() // new

    // Provide db variable to controllers
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    r.GET("/items", controllers.FindItems) // new

    r.Run()
}