package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/zaahidali/education-management-system/handlers"
)

var DB *bun.DB

func main() {
	fmt.Print("\n\nWelcome to the Educational Management System\n\n")

	err := connectToDatabase()
	if err != nil {
		log.Fatal("unable to connect to database")
	}
	fmt.Println("Connected to the database successfully")
	handlers.DB = DB
	fmt.Print("\n\n\n")
	router := gin.New()
	router.GET("/", handlers.HomePage)
	router.GET("/students", handlers.GetStudents)
	router.POST("/students", handlers.AddStudent)
	router.GET("/students/:id", handlers.GetStudentByID)

	router.GET("/courses", handlers.GetCourses)
	router.POST("/courses", handlers.AddNewCourse)
	router.GET("/courses/:id", handlers.GetCourseByCourseId)
	// PUT method is pending

	router.GET("/enrollments", handlers.GetCourseEnrollments)
	router.POST("/enrollments", handlers.AddNewCourseEnrollment)
	router.GET("/enrollments/:id", handlers.GetCourseEnrollmentByID)
	router.GET("/students/:id/courses", handlers.GetCoursesByStudentID) // access all courses per student - each student enrollment to subjects - Get all courses for a student
	router.GET("/courses/:id/students", handlers.GetStudentsByCourseID) // acccess all students by course ID - all enrolled students per subject - Get all students for a course

	// teacher section
	router.Run()
}

func connectToDatabase() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB = bun.NewDB(sqldb, pgdialect.New())
	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return DB.Ping()
}
