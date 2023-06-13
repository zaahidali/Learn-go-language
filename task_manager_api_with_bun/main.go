package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	fmt.Println("Task Manager API Project")


	// Make a connection first
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return
	}
	fmt.Println("Connected to the database", ctx)

	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTask)
	router.DELETE("/tasks/:id", removeTask)
	router.PUT("/tasks/:id", updateTask)
	router.POST("/tasks", addTask)
	router.Run()
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

func updateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
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

func addTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
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