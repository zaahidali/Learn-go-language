package main

import (
	"context"
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
)

type App struct {
	DB *bun.DB
}

func main() {
	fmt.Println("---------- Blog POST ORM System ----------")

	app := &App{}
	err := app.connectToDatabase()
	if err != nil {
		fmt.Println("I am here")
		log.Fatalf("Unable to connect to the database: %s", err.Error())
	}
	fmt.Println("Successfully connected to the database")

	router := gin.New()
	setupRoutes(router, app)
	router.Run()
}

func setupRoutes(router *gin.Engine, app *App) {
	router.GET("/", homePage) // for testing
	// router.GET("/users", app.getUsers)
	router.GET("/users", app.getPaginatedUsers)
	router.POST("/users", app.addNewUser)
	router.GET("/users/:id", app.getUserByID)
	router.DELETE("/users/:id", app.deleteUserByID)

	// posts
	router.GET("/posts", app.getPosts)
	router.POST("/posts", app.addNewPost)
	router.GET("/posts/:id", app.getPostByID)
	router.DELETE("/posts/:id", app.deletePostByID)

	// comments
	router.GET("/comments", app.getComments)
	router.POST("/comments", app.addNewComment)
	router.GET("/comments/:id", app.getCommentByID)
	router.DELETE("/comments/:id", app.deleteCommentByID)

}

func (app *App) connectToDatabase() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_ = context.Background()

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

	app.DB = bun.NewDB(sqldb, pgdialect.New())

	app.DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return app.DB.Ping()
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
