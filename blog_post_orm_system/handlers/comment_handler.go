package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/Learn-go-language/blog_post_orm_system/models"
)

/*
**Comments**
- `POST /comments` - Create a new comment. Accept a JSON body with the comment's content, user_id, and post_id.
- `GET /comments` - Get a list of all comments.
- `GET /comments/:id` - Get the details of a specific comment.
- `PUT /comments/:id` - Update a specific comment. Accept a JSON body with the new details of the comment.
- `DELETE /comments/:id` - Delete a specific comment.
*/

func GetComments(c *gin.Context) {
	var comments []models.Comment

	err := DB.NewSelect().Model(&comments).Order("id ASC").Scan(c.Request.Context())
	if err != nil {
		log.Printf("Error fetching comments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func AddNewComment(c *gin.Context) {
	var addNewComment models.Comment

	if err := c.BindJSON(&addNewComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.NewInsert().Model(&addNewComment).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, addNewComment)
}

func GetCommentByID(c *gin.Context) {
	commentID, valid := ValidateCommentID(c)

	if !valid {
		return
	}

	comment := models.Comment{}
	err := DB.NewSelect().
		Model(&comment).
		Where("id = ?", commentID).
		Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func DeleteCommentByID(c *gin.Context) {
	commentID, valid := ValidateCommentID(c)

	if !valid {
		return
	}

	comment := models.Comment{}
	err := DB.NewDelete().
		Model(&comment).
		Where("id = ?", commentID).
		Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func ValidateCommentID(c *gin.Context) (string, bool) {
	commentID := c.Param("id")
	if commentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment ID is required"})
		return "", false
	}
	return commentID, true
}
