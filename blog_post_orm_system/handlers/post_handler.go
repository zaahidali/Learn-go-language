package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/Learn-go-language/blog_post_orm_system/models"

)

/*
**Posts**
- `POST /posts` - Create a new post. Accept a JSON body with the post's title, content, and user_id.
- `GET /posts` - Get a list of all posts.
- `GET /posts/:id` - Get the details of a specific post.
- `PUT /posts/:id` - Update a specific post. Accept a JSON body with the new details of the post.
- `DELETE /posts/:id` - Delete a specific post.
*/

func GetPosts(c *gin.Context) {
	var posts []models.Post

	err := DB.NewSelect().Model(&posts).Relation("Comment").Order("id ASC").Scan(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No records found"})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})

		return
	}
	c.JSON(http.StatusOK, posts)
}

func AddNewPost(c *gin.Context) {
	var newPost models.Post

	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": newPost})
		return
	}

	if newPost.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	fmt.Printf("New Post: %+v\n", newPost)

	_, err := DB.NewInsert().Model(&newPost).Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "New Post Created"})
}

func GetPostByID(c *gin.Context) {
	postID, valid := ValidatePostID(c)

	if !valid {
		return
	}

	var post models.Post
	err := DB.NewSelect().Model(&post).Relation("Comment").Where("id = ?", postID).Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})

		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePostByID(c *gin.Context) {

	postID, valid := ValidatePostID(c)
	if !valid {
		return
	}

	var post models.Post
	_, err := DB.NewDelete().Model(&post).Where("id = ?", postID).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post successfully deleted!"})
}

func ValidatePostID(c *gin.Context) (string, bool) {
	postID := c.Param("id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be present"})
		return "", false
	}
	return postID, true
}
