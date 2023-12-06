package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"jwt/model" // Adjust this import path according to your project structure
)

var DB *gorm.DB

// Connect initializes a connection to the database
func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Read database connection details from environment variables
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")
	log.Println("user", dbUser, dbPass, dbPort, "user")
	// Construct the database URL
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// Configure Gorm
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Adjust log level as needed
	}

	// Establish a connection to the database
	db, err := gorm.Open(mysql.Open(dbURL), config)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	DB = db

	// AutoMigrate ensures that the User table is created in the database
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Error auto-migrating database: %v", err)
	}

	log.Println("Connected to the database")
}
