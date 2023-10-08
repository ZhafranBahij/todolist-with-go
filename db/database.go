package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error

	//* Turn on database DATAABASE
	dsn := "host=127.0.0.1 user=postgres password= dbname=todolist port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	//* Migrate DATABASE
	// db.AutoMigrate(&User{}, &Todo{})

}
