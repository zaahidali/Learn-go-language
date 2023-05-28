package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllComments(ctx *gin.Context) {
	postID := ctx.Param("id")

	for _, post := range blogPosts {
		if post.ID == postID {
			ctx.JSON(http.StatusOK, post.Comments)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

func addNewPost(ctx *gin.Context) {
	var newPost BlogPost
	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	newPost.ID = generateID()
	blogPosts = append(blogPosts, newPost)
	ctx.JSON(http.StatusCreated, gin.H{"message": "New post added"})
}

func deleteComment(ctx *gin.Context) {
	postID := ctx.Param("id")
	commentID := ctx.Param("commentID")

	for i, post := range blogPosts {
		if post.ID == postID {
			for j, comment := range post.Comments {
				if comment.ID == commentID {
					blogPosts[i].Comments = append(post.Comments[:j], post.Comments[j+1:]...)
					ctx.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
					return
				}
			}
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Post or Comment not found"})
}

func getComment(ctx *gin.Context) {
	postID := ctx.Param("id")
	commentID := ctx.Param("commentID")

	for _, post := range blogPosts {
		if post.ID == postID {
			for _, comment := range post.Comments {
				if comment.ID == commentID {
					ctx.JSON(http.StatusOK, comment)
					return
				}
			}
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Post or Comment not found"})
}

func addNewComment(ctx *gin.Context) {
	postID := ctx.Param("id")

	var newComment Comment
	if err := ctx.ShouldBindJSON(&newComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	for i, post := range blogPosts {
		if post.ID == postID {
			newComment.ID = generateID()
			blogPosts[i].Comments = append(post.Comments, newComment)
			ctx.JSON(http.StatusCreated, newComment)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}
