package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	_ "github.com/joho/godotenv/autoload"     // auto env variables import
)

// GetConnection returns the database connection
func GetConnection() (db *gorm.DB) {
	// Set the database connection
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@/%s",
			os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME")))

	if err != nil {
		log.Fatalln(err)
	}
	// Validate if the connection is right
	if err := db.DB().Ping(); err != nil {
		log.Fatal(err)
	}
	return
}
