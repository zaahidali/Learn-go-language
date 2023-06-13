package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/education-management-system/models"
)

func AddNewCourse(ctx *gin.Context) {
	var course models.Course

	if err := ctx.ShouldBind(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if course.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty"})
		return
	}

	_, err := DB.NewInsert().Model(&course).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert new course", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "New course added successfully"})
}

func GetCourses(ctx *gin.Context) {
	var courses []models.Course

	err := DB.NewSelect().Model(&courses).Order("id ASC").Scan(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "no courses found"})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

func GetCourseByCourseId(ctx *gin.Context) {
	id := ctx.Param("id")
	var course models.Course

	err := DB.NewSelect().Model(&course).Where("id =?", id).Scan(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "no course found"})
		return
	}

	ctx.JSON(http.StatusOK, course)
}
