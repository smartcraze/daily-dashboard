package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/smartcraze/daily-dashboard/db"
	"github.com/smartcraze/daily-dashboard/models"
	"github.com/smartcraze/daily-dashboard/routes"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	fmt.Println("✅ Database connected successfully!")
	// Auto migrate your models
	err = database.AutoMigrate(&models.DailyBrief{}, &models.Task{})

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	app := gin.Default()
	// Register routes
	routes.TaskRoutes(app, database)

	app.Run()

}
