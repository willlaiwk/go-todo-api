package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willlaiwk/go-todo-api/models"
)

// GET /tasks
// Find all tasks
func FindTasks(c *gin.Context) {
	var tasks []models.Task

	models.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// GET /tasks/:id
// Find task by id
func FindTask(c *gin.Context) {
	var task models.Task

	if err := models.DB.First(&task, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
	// Vaildate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	task := models.Task{Title: input.Title, Description: input.Description}
	models.DB.Create(&task)

	c.JSON(http.StatusCreated, gin.H{"data": task})
}

// PATCH /tasks/:id
// Update task
func UpdateTask(c *gin.Context) {
	// Vaildate input
	var patchMap map[string]interface{}

	if err := c.ShouldBindJSON(&patchMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the task exists
	var task models.Task

	if err := models.DB.First(&task, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found."})
		return
	}

	// Update task
	models.DB.Model(&task).Updates(patchMap)

	c.JSON(http.StatusNoContent, nil)
}

// DELETE /tasks/:id
// Delete task by id
func DeleteTask(c *gin.Context) {
	var task models.Task
	if err := models.DB.First(&task, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found."})
		return
	}

	models.DB.Delete(&task)
}
