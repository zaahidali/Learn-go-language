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
	"github.com/zaahidali/Learn-go-language/blog_post_orm_system/handlers"
)

var DB *bun.DB

func main() {
	fmt.Println("---------- Blog POST ORM System ----------")

	err := connectToDatabase()
	if err != nil {
		log.Fatalf("Unable to connect to the database: %s", err.Error())
	}
	fmt.Println("Successfully connected to the database")

	// Assign the database connection to the handlers package's DB variable
	handlers.DB = DB
	router := gin.New()
	setupRoutes(router)
	router.Run()
}

func setupRoutes(router *gin.Engine) {
	router.GET("/", handlers.HomePage) // for testing
	// router.GET("/users", handlers.getUsers)
	router.GET("/users", handlers.GetPaginatedUsers)
	router.POST("/users", handlers.AddNewUser)
	router.GET("/users/:id", handlers.GetUserByID)
	router.DELETE("/users/:id", handlers.DeleteUserByID)

	// posts
	router.GET("/posts", handlers.GetPosts)
	router.POST("/posts", handlers.AddNewPost)
	router.GET("/posts/:id", handlers.GetPostByID)
	router.DELETE("/posts/:id", handlers.DeletePostByID)

	// comments
	router.GET("/comments", handlers.GetComments)
	router.POST("/comments", handlers.AddNewComment)
	router.GET("/comments/:id", handlers.GetCommentByID)
	router.DELETE("/comments/:id", handlers.DeleteCommentByID)
}

func connectToDatabase() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
	// defer sqldb.Close()

	DB = bun.NewDB(sqldb, pgdialect.New())

	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return DB.Ping()
}

/*
**Users**
- `POST /users` - Create a new user. Accept a JSON body with the user's name and email.
- `GET /users` - Get a list of all users.
- `GET /users/:id` - Get the details of a specific user.
- `PUT /users/:id` - Update a specific user. Accept a JSON body with the new details of the user.
- `DELETE /users/:id` - Delete a specific user.

**Posts**
- `POST /posts` - Create a new post. Accept a JSON body with the post's title, content, and user_id.
- `GET /posts` - Get a list of all posts.
- `GET /posts/:id` - Get the details of a specific post.
- `PUT /posts/:id` - Update a specific post. Accept a JSON body with the new details of the post.
- `DELETE /posts/:id` - Delete a specific post.

**Comments**
- `POST /comments` - Create a new comment. Accept a JSON body with the comment's content, user_id, and post_id.
- `GET /comments` - Get a list of all comments.
- `GET /comments/:id` - Get the details of a specific comment.
- `PUT /comments/:id` - Update a specific comment. Accept a JSON body with the new details of the comment.
- `DELETE /comments/:id` - Delete a specific comment.
*/
