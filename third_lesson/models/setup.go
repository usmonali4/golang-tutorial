package models

import (
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbURL := os.Getenv("DB_URL")

	db, err := gorm.Open("postgres", dbURL)

	if err != nil {
		fmt.Println("couldn't connect to db")
		log.Fatalln("connection to db error: ", err)
	} else {
		fmt.Println("We are connected to the database ")
	}

	DB = db

	DB.AutoMigrate(&User{})
}