package controller

import (
	"example/todolist/db"
	"example/todolist/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexUser(c *gin.Context) {

	var users []model.User

	// Get data from database with order desc
	db.DB.Debug().Preload("Todo").Order("id desc").Find(&users)

	c.JSON(200, gin.H{
		"message": http.StatusOK,
		"data":    users,
	})
}

func StoreUser(c *gin.Context) {

	age, _ := strconv.Atoi(c.PostForm("age"))

	new_users := model.User{
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
		Age:   age,
	}

	db.DB.Save(&new_users)

	c.JSON(200, gin.H{
		"message": http.StatusOK,
		"data":    new_users,
	})

}

func UpdateUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	age, _ := strconv.Atoi(c.PostForm("age"))

	user := model.User{
		ID:    uint(id),
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
		Age:   age,
	}

	db.DB.Save(&user)

	c.JSON(200, gin.H{
		"message": http.StatusOK,
		"data":    user,
	})

}

func DeleteUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	db.DB.Debug().Delete(&model.User{}, uint(id))

	c.JSON(200, gin.H{
		"message": http.StatusOK,
	})
}
