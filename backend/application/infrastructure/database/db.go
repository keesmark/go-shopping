package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-shopping/application/model"
)

var (
	db *gorm.DB
	_  error
)

func Init() {
	db = gormConnect()
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	DBNAME := "go"
	CONNECT := USER + ":" + PASS + "@tcp(db:3306)/" + DBNAME + "?parseTime=true"
	database, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return database
}

func autoMigration() {
	db.AutoMigrate(&model.User{})
}
