package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	DBDriver := os.Getenv("DB_DRIVER")
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)

	DB, err = gorm.Open(DBDriver, DBURL)
	if err != nil {
		fmt.Println("Cannot connect to database.")
		log.Fatal("Connection Error:", err)
	} else {
		fmt.Println("Connected to database.")
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Stock{})

}
