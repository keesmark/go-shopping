package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go-shopping/application"
	"go-shopping/application/infrastructure/database"
)

func main() {
	database.Init()
	application.InitRouter()
}
