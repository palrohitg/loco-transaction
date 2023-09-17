package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"loco-transaction/models"
	"log"
	"os"
)

var Con *gorm.DB

func Connect() *gorm.DB {
	fmt.Println("Starting the container here")
	if Con != nil {
		return Con
	}
	fmt.Println("Starting the container here 2")
	dbURL := os.Getenv("LOCO_DATABASE_URL")
	fmt.Println(dbURL)
	if dbURL == "" {
		db, err := gorm.Open(mysql.Open("tester:secret@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
			log.Fatalf("Got error when connect database, the error is '%v'", err)
		}
		db.AutoMigrate(&models.TransactionSummary{})
		Con = db
		return db
	} else {
		db, err := gorm.Open(mysql.Open("tester:secret@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Got error jwhen connect database, the error is '%v'", err)
		}
		db.AutoMigrate(&models.TransactionSummary{})
		Con = db
		return db
	}
}
