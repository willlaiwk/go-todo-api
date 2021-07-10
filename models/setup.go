package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "todo:todo123@tcp(127.0.0.1:33306)/todo?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Task{})

	DB = db
}
