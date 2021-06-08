package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-shopping/db"
	"go-shopping/models"
	"net/http"
)

func CreateTodoData(c *gin.Context) {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		fmt.Errorf("%#v", err)
	}
	todo.Status = 0

	fmt.Println(todo)
	database := db.GetDB()
	//defer database.Close()
	database.Create(&todo)

	todos := getAllTodos()
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todos": todos,
	})
}

func UpdateDoneTodoData(c *gin.Context) {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		fmt.Errorf("%#v", err)
	}

	database := db.GetDB()
	result := database.First(&todo, todo.Id)
	fmt.Println(result)
	if todo.Status == 0 {
		todo.Status = 1
	} else {
		todo.Status = 0
	}
	database.Save(&todo)
	//db.Close()
	todos := getAllTodos()
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todos": todos,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./template/index.html")
	db.Init()

	todos := getAllTodos()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"todos": todos,
		})
	})

	r.POST("/todo", CreateTodoData)
	r.POST("/done", UpdateDoneTodoData)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//db.Close()
}

func getAllTodos() []models.Todo {
	database := db.GetDB()
	//defer database.Close()
	var todos []models.Todo
	database.Order("created_at desc").Find(&todos)
	return todos
}
