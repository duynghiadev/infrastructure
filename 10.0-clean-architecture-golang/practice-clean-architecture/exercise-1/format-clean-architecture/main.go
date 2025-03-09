package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	todotrpt "first-app/module/item/transport"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	mysqlConnStr := os.Getenv("MYSQL_CONNECTION")
	if mysqlConnStr == "" {
		log.Fatalln("Missing MySQL connection string. Please set MYSQL_CONNECTION environment variable.")
	}

	log.Println("MYSQL_CONNECTION:", mysqlConnStr)

	// Connect to MySQL
	db, err := gorm.Open(mysql.Open(mysqlConnStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to MySQL: %v\n", err)
	}

	log.Println("Connected to MySQL successfully.")

	// Set up Gin router
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/items", todotrpt.HandleCreateItem(db))         // create item
		v1.GET("/items", todotrpt.HandleListItem(db))            // list items
		v1.GET("/items/:id", todotrpt.HandleFindAnItem(db))      // get an item by ID
		v1.PUT("/items/:id", todotrpt.HandleUpdateAnItem(db))    // edit an item by ID
		v1.DELETE("/items/:id", todotrpt.HandleDeleteAnItem(db)) // delete an item by ID
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server is running on port %s...\n", port)
	router.Run(":" + port)
}
