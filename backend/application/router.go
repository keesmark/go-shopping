package application

import (
	"github.com/gin-gonic/gin"
	"go-shopping/application/controller"
)

func InitRouter() {
	r := gin.Default()

	todoController := controller.TodoController{}
	r.GET("/", func(c *gin.Context) {
		todoController.GetAllTodos(c)
	})
	r.POST("/todo/store", func(c *gin.Context) {
		todoController.CreateTodo(c)
	})
	r.POST("/status/:id", func(c *gin.Context) {
		todoController.UpdateDoneTodoData(c)
	})
	r.POST("/delete/:id", func(c *gin.Context) {
		todoController.DeleteTodo(c)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
