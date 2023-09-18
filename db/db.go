package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"loco-transaction/models"
	"log"
	"os"
)

var Con *gorm.DB

func Connect() *gorm.DB {
	if Con != nil {
		return Con
	}
	dbURL := os.Getenv(LocoDatabaseUrlkey)
	if dbURL == "" {
		db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("Got error when connect database, the error is '%v'", err)
		}
		db.AutoMigrate(&models.TransactionSummary{})
		Con = db
		return db
	} else {
		db, err := gorm.Open(mysql.Open("sergey:sergey@tcp(db:3306)/local_transaction?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Got error jwhen connect database, the error is '%v'", err)
		}
		db.AutoMigrate(&models.TransactionSummary{})
		Con = db
		return db
	}
}
