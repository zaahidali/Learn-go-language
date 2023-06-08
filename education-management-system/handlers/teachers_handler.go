package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/education-management-system/models"
)

func GetTeachers(ctx *gin.Context) {
	var teachers []models.Teacher

	err := DB.NewSelect().Model(&teachers).Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, teachers)
}

func GetTeachersByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var teacher models.Teacher

	err := DB.NewSelect().Model(&teacher).Where("id =?", id).Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, teacher)
}

func AddNewCourseTeacher(ctx *gin.Context) {
	var teacher models.Teacher

	err := ctx.ShouldBind(&teacher)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = DB.NewInsert().Model(&teacher).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, teacher)
}

func RemoveTeachByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var teacher models.Teacher

	_, err := DB.NewDelete().Model(&teacher).Where("id =?", id).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Teacher removed"})
}

func UpdateTeacherByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var teacher models.Teacher

	err := ctx.ShouldBind(&teacher)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = DB.NewUpdate().Model(&teacher).Where("id =?", id).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Teacher updated"})
}
