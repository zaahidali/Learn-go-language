package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/zaahidali/task_manager_api_with_bun/model" // we imported the model here
)

var DB *bun.DB

func HomePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task Manager API"})
}

func GetTasks(ctx *gin.Context) {
	// Create a slice to store the retrieved tasks
	var tasks []model.Task

	// Execute the database query to retrieve tasks using Go bun
	err := DB.NewSelect().Model(&tasks).Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the retrieved tasks in the response
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTask(c *gin.Context) {
	taskID := c.Param("id")

	// Check if the task ID is empty
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be present"})
		return
	}

	task := &model.Task{}

	// Fetch specific record from the database using Go bun
	err := DB.NewSelect().Model(task).Where("id = ?", taskID).Scan(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if task.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(ctx *gin.Context) {
	taskID := ctx.Param("id")

	if taskID == "" {
		ctx.JSON(http.StatusNoContent, gin.H{"error": "ID must be present"})
		return
	}

	updatedTask := &model.Task{}

	// Bind JSON body to the updatedTask struct
	if err := ctx.ShouldBindJSON(updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the task record in the database using Go bun
	_, err := DB.NewUpdate().Model(updatedTask).
		Set("title = ?", updatedTask.Title).
		Set("description = ?", updatedTask.Description).
		Where("id = ?", taskID).
		Exec(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func RemoveTask(ctx *gin.Context) {
	taskID := ctx.Param("id")

	task := &model.Task{}

	// Delete specific task record from the database
	res, err := DB.NewDelete().Model(task).Where("id = ?", taskID).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if any rows were affected by the delete operation
	if rowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
}

func AddTask(ctx *gin.Context) {
	newTask := &model.Task{}

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the new task record into the database
	_, err := DB.NewInsert().Model(newTask).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}
