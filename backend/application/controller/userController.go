package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-shopping/application/infrastructure/database"
	"go-shopping/application/model"
	"net/http"
)

type UserController struct {
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		_ = fmt.Errorf("%#v", err)
	}

	db := database.GetDB()
	db.Create(&user)

	c.JSON(http.StatusCreated, map[string]interface{}{
		"user": user,
	})
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	db := database.GetDB()
	db.Find(&user, id)

	db.Save(&user)

	c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "ok",
	})
}
func (controller *UserController) GetUsers(c *gin.Context) {
	db := database.GetDB()
	var users []model.User
	db.Order("created_at desc").Find(&users)

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": users,
	})
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user model.User
	db := database.GetDB()
	db.Delete(&user, id)
	c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "ok",
	})
}
