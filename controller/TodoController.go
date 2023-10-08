package controller

import (
	"example/todolist/db"
	"example/todolist/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexTodo(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("userid"))

	var todo []model.Todo

	db.DB.Debug().Order("id desc").Where("user_id = ?", uint(id)).Find(&todo)

	c.JSON(200, gin.H{
		"message": http.StatusOK,
		"data":    todo,
	})
}

func StoreTodo(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("userid"))
	event := c.PostForm("event")

	todo := model.Todo{
		UserID: uint(id),
		Event:  event,
	}

	db.DB.Debug().Save(&todo)

	c.JSON(200, gin.H{
		"message": http.StatusOK,
		"data":    todo,
	})
}

func UpdateTodo(c *gin.Context) {

	// Get data from PARAM
	id, _ := strconv.Atoi(c.Param("userid"))
	todo_id, _ := strconv.Atoi(c.Param("todoid"))

	// Input from Form
	event := c.PostForm("event")

	todo := model.Todo{
		ID:     uint(todo_id),
		UserID: uint(id),
		Event:  event,
	}

	// Save data to db
	db.DB.Debug().Save(&todo)

	c.JSON(200, gin.H{
		"message": http.StatusOK,
		"data":    todo,
	})
}

func DeleteTodo(c *gin.Context) {

	// get data from parameter
	todo_id, _ := strconv.Atoi(c.Param("todoid"))

	// delete data with id
	db.DB.Delete(&model.Todo{}, uint(todo_id))

	c.JSON(200, gin.H{
		"message": http.StatusOK,
	})
}
