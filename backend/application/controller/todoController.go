package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-shopping/application/infrastructure/database"
	"go-shopping/application/model"
	"net/http"
)

type TodoController struct {
}

func (controller *TodoController) CreateTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		_ = fmt.Errorf("%#v", err)
	}
	todo.Status = 0

	db := database.GetDB()
	db.Create(&todo)

	c.JSON(http.StatusCreated, map[string]interface{}{
		"todo": todo,
	})
}

func (controller *TodoController) UpdateDoneTodoData(c *gin.Context) {
	var todo model.Todo
	id := c.Param("id")
	db := database.GetDB()
	db.Find(&todo, id)

	if todo.Status == 0 {
		todo.Status = 1
	} else {
		todo.Status = 0
	}
	db.Save(&todo)

	c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "ok",
	})
}
func (controller *TodoController) GetAllTodos(c *gin.Context) {
	db := database.GetDB()
	var todos []model.Todo
	db.Order("created_at desc").Find(&todos)

	c.JSON(http.StatusOK, map[string]interface{}{
		"todo": todos,
	})
}

func (controller *TodoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	var todo model.Todo
	db := database.GetDB()
	//db.First(&todo, id)
	db.Delete(&todo, id)
	c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "ok",
	})
}
