package main

import (
	"fmt"
	"go-clean-translation/controller/httpapi"
	"go-clean-translation/infras/googlesv"
	mysqlRepo "go-clean-translation/infras/mysql"
	translateServ "go-clean-translation/service"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := connectDBWithRetry(5)

	if err != nil {
		log.Fatalf("[error] Failed to initialize database: %v\n", err)
	}

	// Setup Dependencies
	repository := mysqlRepo.NewMySQLRepo(db)
	googleTranslate := googlesv.New()
	service := translateServ.NewService(repository, googleTranslate)
	controller := httpapi.NewAPIController(service)

	engine := gin.Default()

	v1 := engine.Group("/v1")
	controller.SetUpRoute(v1)

	log.Println("Server is running on port 8080...")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("[error] Server failed to start: %v\n", err)
	}
}

func connectDBWithRetry(times int) (*gorm.DB, error) {
	var e error
	defaultDSN := "root:duynghia123@tcp(127.0.0.1:3306)/demo_db?charset=utf8mb4&parseTime=True&loc=Local"

	for i := 1; i <= times; i++ {
		dsn := os.Getenv("MYSQL_DSN")
		if dsn == "" {
			log.Println("[warning] MYSQL_DSN is empty, using default DSN")
			dsn = defaultDSN
		}

		fmt.Printf("[info] Attempting to connect to DB (try %d/%d): %s\n", i, times, dsn)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err == nil {
			log.Println("[success] Connected to database successfully")
			return db, nil
		}

		e = err
		log.Printf("[error] Failed to connect to DB: %v. Retrying in 2s...\n", err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", times, e)
}
