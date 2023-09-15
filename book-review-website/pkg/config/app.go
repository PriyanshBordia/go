package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

var username = os.Getenv("username")
var password = os.Getenv("password")
var DB = os.Getenv("db")


func Connect() {
	conn, err := gorm.Open("mysql", username + ":" + password + "@tcp(localhost:3306)/" + DB + "?charset=utf8&parseTime=True")
	// db, err = gorm.Open("sqlite3", "test.db")
	// db, err = gorm.Open("postgres", "user=" + username + "password=" + password + DB.name=" + DB + " port=9920 sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}