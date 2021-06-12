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
	database.Create(&todo)

	c.Redirect(http.StatusSeeOther, "/todo")
}

func UpdateDoneTodoData(c *gin.Context) {
	var todo models.Todo
	id := c.Param("Id")
	database := db.GetDB()
	database.Find(&todo, id)

	if todo.Status == 0 {
		todo.Status = 1
	} else {
		todo.Status = 0
	}
	database.Save(&todo)

	c.Redirect(http.StatusSeeOther, "/todo")
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./template/index.html")
	db.Init()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/todo", getAllTodos)
	r.POST("/todo/store", CreateTodoData)
	r.POST("/status/:Id", UpdateDoneTodoData)
	r.POST("/delete/:Id", deleteTodo)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getAllTodos(c *gin.Context) {
	database := db.GetDB()
	var todos []models.Todo
	database.Order("created_at desc").Find(&todos)

	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todos": todos,
	})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("Id")

	var todo models.Todo
	database := db.GetDB()
	database.First(&todo, id)
	database.Delete(&todo)
	c.Redirect(http.StatusSeeOther, "/todo")
}
