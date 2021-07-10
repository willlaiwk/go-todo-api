package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willlaiwk/go-todo-api/api"
	"github.com/willlaiwk/go-todo-api/models"
)

func main() {
	r := gin.Default()

	// setup api routes
	api.InitAPIRoutes(r)

	// setup and connect database
	models.ConnectDatabase()

	// start server
	r.Run()
}
