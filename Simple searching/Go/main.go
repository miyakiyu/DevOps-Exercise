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
	dsn := "root:12345@tcp(127.0.0.1:3306)/Cats?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := gin.Default()

	r.GET("/cats", func(c *gin.Context) {
		cats, err := models.GetAllCats(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cats)
	})

	r.Run(":8080")
}
