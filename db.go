package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "sql12828936:aF8eADc2mG@tcp(sql12.freesqldatabase.com:3306)/sql12828936?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	DB = db
	log.Println("Connected to MySQL with GORM")
}
