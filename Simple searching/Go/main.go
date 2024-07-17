package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"main/models"
)

func main() {
	// Globla setting
	dsn := "root:12345@tcp(172.24.0.1:3306)/cat"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := gin.Default()
	// Home Page
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	// Get cats
	r.GET("/cats", func(c *gin.Context) {
		cats, err := models.GetCats(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cats)
	})

	r.Run(":8080")
}
