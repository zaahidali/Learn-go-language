package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Simple Blog"})
}

func getPosts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, blogPosts)
}

func getPost(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, post := range blogPosts {
		if post.ID == id {
			ctx.JSON(http.StatusOK, post)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

func deletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, post := range blogPosts {
		if post.ID == id {
			blogPosts = append(blogPosts[:i], blogPosts[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Post removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

func updatePost(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedPost BlogPost
	if err := ctx.ShouldBindJSON(&updatedPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	for i, post := range blogPosts {
		if post.ID == id {
			if updatedPost.Title != "" {
				blogPosts[i].Title = updatedPost.Title
			}
			if updatedPost.Content != "" {
				blogPosts[i].Content = updatedPost.Content
			}
			if updatedPost.Author != "" {
				blogPosts[i].Author = updatedPost.Author
			}

			ctx.JSON(http.StatusOK, gin.H{"message": "Post updated"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}
