package main

import (
	"crowdfounding/user"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/users", userHandler)
	router.Run()
}

func userHandler(c *gin.Context) {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	var users []user.User

	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
