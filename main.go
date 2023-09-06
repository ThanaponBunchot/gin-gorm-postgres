package main

import (
	"gin-gorm-postgres/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

// func init() {

// 	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5435 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return
// 	}
// 	db.AutoMigrate(&model.User{}, &model.CreditCard{})
// }

func main() {
	router := gin.Default()

	// Define an API route for creating a user
	router.POST("/create-user", createUser)

	// Start the Gin server
	router.Run(":8080")
	// Migrate the schema

}

func createUser(c *gin.Context) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5435 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&model.User{}, &model.CreditCard{})

	// Create a user with associated credit cards
	user := model.User{
		CreditCards: []model.CreditCard{
			{Number: "dd", UserID: 1},
			{Number: "dd", UserID: 2},
		},
	}

	// Save the user and associated credit cards to the database
	if err := db.Create(&user).Error; err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
