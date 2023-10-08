package main

import (
	"example/todolist/controller"
	"example/todolist/db"

	"github.com/gin-gonic/gin"
)

func main() {

	//* Connect to DB
	db.ConnectDB()

	router := gin.Default()

	//* GET ALL USERS
	router.GET("/users", controller.IndexUser)

	//* CREATE USERS
	router.POST("/users", controller.StoreUser)

	//* UPDATE USERS
	router.PUT("/users/:id", controller.UpdateUser)

	//* DELETE USERS
	router.DELETE("/users/:id", controller.DeleteUser)

	//* GET TODOLIST in THAT USER
	router.GET("/users/:userid/todolist", controller.IndexTodo)

	//* POST TODOLIST in THAT USER
	router.POST("/users/:userid/todolist", controller.StoreTodo)

	//* PUT TODOLIST in THAT USER
	// router.PUT("/users/:userid/todolist/:todoid", controller.UpdateTodo)

	//* DELETE USERS
	router.DELETE("/users/:id/todolist/:todoid", controller.DeleteTodo)

	router.Run() // listen and serve on 0.0.0.0:8080

}
