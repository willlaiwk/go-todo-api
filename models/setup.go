package models

import (
	"fmt"

	"github.com/willlaiwk/go-todo-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: connectionString,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Task{})

	DB = db
}
