package main

import (
	"github.com/PacoDw/introduction_to_GORM/db"
	"github.com/PacoDw/introduction_to_GORM/models"
	"github.com/PacoDw/introduction_to_GORM/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	// Get the database connection
	db := db.GetConnection()
	defer db.Close()

	// Enable the log showing detailed log
	db.LogMode(true)

	// Migrate the Models to the database
	// db.Debug().DropTableIfExists(&models.User{}) // drop db table
	db.Debug().AutoMigrate(&models.User{})
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	// router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the user
	routes.User(router.Group("/user"))

	router.Run(":5000")
}
