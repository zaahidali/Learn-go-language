package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
**Comments**
- `POST /comments` - Create a new comment. Accept a JSON body with the comment's content, user_id, and post_id.
- `GET /comments` - Get a list of all comments.
- `GET /comments/:id` - Get the details of a specific comment.
- `PUT /comments/:id` - Update a specific comment. Accept a JSON body with the new details of the comment.
- `DELETE /comments/:id` - Delete a specific comment.
*/

func (app *App) getComments(c *gin.Context) {
	var comments []Comment

	err := app.DB.NewSelect().Model(&comments).Order("id ASC").Scan(c.Request.Context())
	if err != nil {
		log.Printf("Error fetching comments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (app *App) addNewComment(c *gin.Context) {
	var addNewComment Comment

	if err := c.BindJSON(&addNewComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := app.DB.NewInsert().Model(&addNewComment).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addNewComment)
}

func (app *App) getCommentByID(c *gin.Context) {
	commentID, valid := validateCommentID(c)

	if !valid {
		return
	}

	comment := Comment{}
	err := app.DB.NewSelect().
		Model(&comment).
		Where("id = ?", commentID).
		Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (app *App) deleteCommentByID(c *gin.Context) {
	commentID, valid := validateCommentID(c)

	if !valid {
		return
	}

	comment := Comment{}
	err := app.DB.NewDelete().
		Model(&comment).
		Where("id = ?", commentID).
		Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func validateCommentID(c *gin.Context) (string, bool) {
	commentID := c.Param("id")
	if commentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment ID is required"})
		return "", false
	}
	return commentID, true
}
