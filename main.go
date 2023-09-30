package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todo struct {
	ID        uint
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
	Event     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {

	//* Turn on database DATAABASE
	dsn := "host=127.0.0.1 user=postgres password= dbname=todolist port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//* Migrate DATABASE
	// db.AutoMigrate(&User{}, &Todo{})

	router := gin.Default()

	//* GET ALL USERS
	router.GET("/users", func(c *gin.Context) {

		var users []User

		// Get data from database with order desc
		db.Debug().Preload("Todo").Order("id desc").Find(&users)

		c.JSON(200, gin.H{
			"message": http.StatusOK,
			"data":    users,
		})
	})

	//* CREATE USERS
	router.POST("/users", func(c *gin.Context) {

		age, _ := strconv.Atoi(c.PostForm("age"))

		new_users := User{
			Name:  c.PostForm("name"),
			Email: c.PostForm("email"),
			Age:   age,
		}

		db.Save(&new_users)

		c.JSON(200, gin.H{
			"message": http.StatusOK,
			"data":    new_users,
		})

	})

	//* UPDATE USERS
	router.PUT("/users/:id", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))
		age, _ := strconv.Atoi(c.PostForm("age"))

		user := User{
			ID:    uint(id),
			Name:  c.PostForm("name"),
			Email: c.PostForm("email"),
			Age:   age,
		}

		db.Save(&user)

		c.JSON(200, gin.H{
			"message": http.StatusOK,
			"data":    user,
		})

	})

	//* DELETE USERS
	router.DELETE("/users/:id", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))

		db.Delete(&User{}, uint(id))

		c.JSON(200, gin.H{
			"message": http.StatusOK,
		})
	})

	//* GET TODOLIST in THAT USER
	router.GET("/users/:userid/todolist", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("userid"))

		var todo []Todo

		db.Debug().Order("id desc").Where("user_id = ?", uint(id)).Find(&todo)

		c.JSON(200, gin.H{
			"message": http.StatusOK,
			"data":    todo,
		})
	})

	//* POST TODOLIST in THAT USER
	router.POST("/users/:userid/todolist", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("userid"))
		event := c.PostForm("event")

		todo := Todo{
			UserID: uint(id),
			Event:  event,
		}

		db.Debug().Save(&todo)

		c.JSON(200, gin.H{
			"message": http.StatusOK,
			"data":    todo,
		})
	})

	//* PUT TODOLIST in THAT USER
	router.PUT("/users/:userid/todolist/:todoid", func(c *gin.Context) {

		// Get data from PARAM
		id, _ := strconv.Atoi(c.Param("userid"))
		todo_id, _ := strconv.Atoi(c.Param("todoid"))

		// Input from Form
		event := c.PostForm("event")

		todo := Todo{
			ID:     uint(todo_id),
			UserID: uint(id),
			Event:  event,
		}

		// Save data to db
		db.Debug().Save(&todo)

		c.JSON(200, gin.H{
			"message": http.StatusOK,
			"data":    todo,
		})
	})

	//* DELETE USERS
	router.DELETE("/users/:id/todolist/:todoid", func(c *gin.Context) {

		// get data from parameter
		todo_id, _ := strconv.Atoi(c.Param("todoid"))

		// delete data with id
		db.Delete(&Todo{}, uint(todo_id))

		c.JSON(200, gin.H{
			"message": http.StatusOK,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080

}
