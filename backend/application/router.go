package application

import (
	"github.com/gin-gonic/gin"
	"go-shopping/application/controller"
)

func InitRouter() {
	r := gin.Default()

	userController := controller.UserController{}
	api := r.Group("/api")
	{
		user := api.Group("/users")
		{
			user.GET("/", func(c *gin.Context) {
				userController.GetUsers(c)
			})
			user.POST("/create", func(c *gin.Context) {
				userController.CreateUser(c)
			})
			user.POST("/:id/update", func(c *gin.Context) {
				userController.UpdateUser(c)
			})
			user.POST("/:id/delete", func(c *gin.Context) {
				userController.DeleteUser(c)
			})
		}
	}
	err := r.Run()
	if err != nil {
		return
	}
}
