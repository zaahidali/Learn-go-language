package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
The API supports the following operations:

- GET /tasks: Get a list of all tasks.
- GET /tasks/:id: Get the details of a specific task.
- PUT /tasks/:id: Update a specific task. This should accept a JSON body with the new details of the task.
- DELETE /tasks/:id: Delete a specific task.
- POST /tasks: Create a new task. This should accept a JSON body with the task's title, description, due date, and status.
*/

func main() {
	fmt.Println("Task Manager API Project")

	router := gin.Default()

	router.GET("/", homePage)
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTask)
	router.DELETE("/tasks/:id", removeTask)
	router.POST("/tasks", addTask)
	router.PUT("/tasks/:id", updateTask)
	router.Run()
}

func homePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task Manager API"})
}

func getTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func getTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, val := range tasks {
		if val.ID == id {
			ctx.JSON(http.StatusOK, val)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func removeTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, val := range tasks {
		if val.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func addTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func updateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			// Update only the specified fields
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task Manager Project", Description: "Add/View/Delete Tasks", DueDate: time.Now(), Status: "In Progress"},
	{ID: "2", Title: "Books Management Project", Description: "Add/View/Delete Books", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},
}