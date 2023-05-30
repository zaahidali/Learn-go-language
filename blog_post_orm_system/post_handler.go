package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
**Posts**
- `POST /posts` - Create a new post. Accept a JSON body with the post's title, content, and user_id.
- `GET /posts` - Get a list of all posts.
- `GET /posts/:id` - Get the details of a specific post.
- `PUT /posts/:id` - Update a specific post. Accept a JSON body with the new details of the post.
- `DELETE /posts/:id` - Delete a specific post.
*/

func (app *App) getPosts(c *gin.Context) {
	var posts []Post

	err := app.DB.NewSelect().Model(&posts).Relation("Comment").Order("id ASC").Scan(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No records found"})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})

		return
	}
	c.JSON(http.StatusOK, posts)
}

func (app *App) addNewPost(c *gin.Context) {
	var newPost Post

	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": newPost})
		return
	}

	if newPost.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	fmt.Printf("New Post: %+v\n", newPost)

	_, err := app.DB.NewInsert().Model(&newPost).Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "New Post Created"})
}

func (app *App) getPostByID(c *gin.Context) {
	postID, valid := validatePostID(c)

	if !valid {
		return
	}

	var post Post
	err := app.DB.NewSelect().Model(&post).Relation("Comment").Where("id = ?", postID).Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})

		return
	}

	c.JSON(http.StatusOK, post)
}

func (app *App) deletePostByID(c *gin.Context) {

	postID, valid := validatePostID(c)
	if !valid {
		return
	}

	var post Post
	_, err := app.DB.NewDelete().Model(&post).Where("id = ?", postID).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post successfully deleted!"})
}

func validatePostID(c *gin.Context) (string, bool) {
	postID := c.Param("id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be present"})
		return "", false
	}
	return postID, true
}
