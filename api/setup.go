package api

import "github.com/gin-gonic/gin"

func InitAPIRoutes(r *gin.Engine) {
	r.GET("/api/tasks", FindTasks)
	r.GET("/api/tasks/:id", FindTask)
	r.POST("/api/tasks", CreateTask)
	r.PATCH("/api/tasks/:id", UpdateTask)
	r.DELETE("/api/tasks/:id", DeleteTask)
}
