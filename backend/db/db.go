package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-shopping/models"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db := gormConnect()
	defer db.Close()
	db.AutoMigrate(&models.Todo{})
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
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
