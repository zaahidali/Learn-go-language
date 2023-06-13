package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/zaahidali/education-management-system/models"
)

var DB *bun.DB

func HomePage(ctx *gin.Context) {
	ctx.JSON(http.StatusFound, gin.H{"message": "This is testing home page"})
}

func GetStudents(ctx *gin.Context) {
	var students []models.Student

	err := DB.NewSelect().Model(&students).Order("id ASC").Scan(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No records found"})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func AddStudent(ctx *gin.Context) {
	var student models.Student

	if err := ctx.ShouldBind(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Print("\n\nStudent data:", student, "\n\n")
	_, err := DB.NewInsert().Model(&student).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, student)
}

func GetStudentByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var student models.Student

	err := DB.NewSelect().Model(&student).Where("id =?", id).Scan(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No records found"})
		return
	}

	ctx.JSON(http.StatusOK, student)
}
