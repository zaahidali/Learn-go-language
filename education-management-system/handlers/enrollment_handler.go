package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/education-management-system/models"
)

func AddNewCourseEnrollment(ctx *gin.Context) {
	var newEnrollment models.Enrollment

	if err := ctx.ShouldBind(&newEnrollment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.NewInsert().Model(&newEnrollment).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to insert new data", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "new enrollment added"})
}

func GetCourseEnrollments(ctx *gin.Context) {
	var enrollments []models.Enrollment

	err := DB.NewSelect().Model(&enrollments).Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve data", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, enrollments)
}

func GetCourseEnrollmentByID(ctx *gin.Context) {
	var enrollment models.Enrollment

	err := DB.NewSelect().Model(&enrollment).Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve data", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, enrollment)
}

func GetCoursesByStudentID(ctx *gin.Context) {
	studentID := ctx.Param("id")
	var enrollments []models.Enrollment

	err := DB.NewSelect().Model(&enrollments).Where("student_id = ?", studentID).Relation("Course").Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve data", "error": err.Error()})
		return
	}

	courses := make([]models.Course, len(enrollments)) // fetch all courses for each enrollments
	for i, enrollment := range enrollments {
		courses[i] = *enrollment.Course
	}

	ctx.JSON(http.StatusOK, courses)
}

func GetStudentsByCourseID(ctx *gin.Context) {
	courseID := ctx.Param("id")

	var enrollments []models.Enrollment
	err := DB.NewSelect().Model(&enrollments).Where("course_id = ?", courseID).Relation("Student").Scan(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve data", "error": err.Error()})
		return
	}

	students := make([]models.Student, len(enrollments))
	for i, enrollmentenrollments := range enrollments {
		students[i] = *enrollmentenrollments.Student
	}

	ctx.JSON(http.StatusOK, students)
}

// func GetCoursesByStudentID(ctx *gin.Context) {
// 	studentID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
// 		return
// 	}

// 	var enrollments []models.Enrollment
// 	err = DB.NewSelect().
// 		Model(&enrollments).
// 		Where("student_id = ?", studentID).
// 		Relation("Course").
// 		Scan(ctx.Request.Context())
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve data", "error": err.Error()})
// 		return
// 	}

// 	courses := make([]models.Course, len(enrollments))
// 	for i, enrollment := range enrollments {
// 		courses[i] = *enrollment.Course
// 	}

// 	ctx.JSON(http.StatusOK, courses)
// }
