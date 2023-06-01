package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/zaahidali/Learn-go-language/blog_post_orm_system/models"
)

var DB *bun.DB

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Blog POSTS ORM API"})
}

// get all users - instead of paginated
// func getUsers(c *gin.Context) {
// 	var users []User

// 	err := DB.NewSelect().Model(&users).Relation("Post").Order("id ASC").Scan(c.Request.Context())

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, users)
// }

func AddNewUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if newUser.Name == "" || newUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Email are required fields"})
		return
	}

	_, err := DB.NewInsert().Model(&newUser).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func GetPaginatedUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageLimit, err := strconv.Atoi(page)
	if err != nil || pageLimit <= 0 {
		pageLimit = 1
	}
	limitNumber, err := strconv.Atoi(limit)
	if err != nil || limitNumber <= 0 {
		limitNumber = 10
	}
	offset := (pageLimit - 1) * limitNumber

	var users []models.User

	err = DB.NewSelect().
		Model(&users).Limit(limitNumber).
		Offset(offset).Order("id ASC").Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No data found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"page":  pageLimit,
		"limit": limitNumber,
		"users": users})
}

func GetUserByID(c *gin.Context) {
	userID, valid := ValidateUserID(c)

	if !valid {
		return
	}

	var user models.User

	err := DB.NewSelect().
		Model(&user).Relation("Post").Where("id = ?", userID).Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No record found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUserByID(c *gin.Context) {
	userID, valid := ValidateUserID(c)

	if !valid {
		return
	}

	var user models.User
	_, err := DB.NewDelete().Model(&user).Where("id = ?", userID).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No record found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted Successfully"})
}

func ValidateUserID(c *gin.Context) (string, bool) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be present"})
		return "", false
	}
	return userID, true
}
