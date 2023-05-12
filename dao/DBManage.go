package dao

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var db_address string
var db_username string
var db_secret string

func init() {
	db_address = os.Getenv("DATABASE_ADDRESS")
	db_username = os.Getenv("DATABASE_USERNAME")
	db_secret = os.Getenv("DATABASE_PASSWORD")

	dsn := db_username + ":" + db_secret + "@tcp(" + db_address + ")/msdm?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // DSN data source name
		DefaultStringSize: 256}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// AutoMigrate()
}

func AutoMigrate() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Deck{})
	db.AutoMigrate(&Collection{})
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&Variant{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Artist{})
	db.AutoMigrate(&Title{})
	db.AutoMigrate(&Location{})
}
