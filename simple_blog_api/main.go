package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*
The API supports the following operations:
1. POST /posts - Create a new blog post. This endpoint accepts a JSON body with the post's title, content, and author.
2. GET /posts - Get a list of all blog posts.
3. GET /posts/:id - Get the details of a specific blog post.
4. PUT /posts/:id - Update a specific blog post. This endpoint accepts a JSON body with the new details of the post.
5. DELETE /posts/:id - Delete a specific blog post.

Additionally, the API also provides operations for managing comments:
1. POST /posts/:id/comments - Add a new comment to a specific blog post. This endpoint accepts a JSON body with the comment's content and author.
2. GET /posts/:id/comments - Get all comments for a specific blog post.
3. DELETE /posts/:id/comments/:commentId - Delete a specific comment from a specific blog post.
*/

func main() {
	fmt.Println("Simple Blog Api")
	router := gin.New()

	router.GET("/", homePage)
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPost)
	router.PUT("/posts/:id", updatePost)
	router.DELETE("/posts/:id", deletePost)
	router.POST("/posts", addNewPost)

	// comments
	// get all comments for a particular post
	router.GET("/posts/:id/comments", getAllComments)
	router.GET("/posts/:id/comments/:commentID", getComment)
	router.POST("/posts/:id/comments", addNewComment)
	router.DELETE("/posts/:id/comments/:commentID", deleteComment)
	router.Run()
}

func generateID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}


var blogPosts = []BlogPost{
	{
		ID:      "1",
		Title:   "Learn Go lang",
		Content: "This is a blog post about learning Go programming language. It includes tips and resources for beginners and advanced programmers alike.",
		Author:  "John Doe",
		Comments: []Comment{
			{
				ID:      "1",
				Content: "Great post! I found it very helpful.",
				Author:  "Commenter One",
			},
			{
				ID:      "2",
				Content: "I agree, this is a very informative post.",
				Author:  "Commenter Two",
			},
		},
	},
	{
		ID:      "2",
		Title:   "Intro to gin-gonic",
		Content: "This post is about gin-gonic, a web framework for Go that's great for building REST APIs. We'll go over the basics and show some examples of how to use it.",
		Author:  "Jane Smith",
		Comments: []Comment{
			{
				ID:      "3",
				Content: "Great post! I found it very helpful.",
				Author:  "Commenter Three",
			},
			{
				ID:      "4",
				Content: "I agree, this is a very informative post.",
				Author:  "Commenter Four",
			},
		},
	},
}
