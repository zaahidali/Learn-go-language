package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func welcomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Budget Management System"})
}

func addNewItem(c *gin.Context) {
	var newItem Budget

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err := db.NewInsert().Model(&newItem).Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newItem)
}

// func getAllItems(c *gin.Context) {
// 	var allItems []Budget
// 	err := db.NewSelect().Model(&allItems).Order("id ASC").Scan(c.Request.Context())
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, allItems)
// }

func getPaginatedItems(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageLimit, err := strconv.Atoi(page)
	if err != nil || pageLimit <= 0 {
		pageLimit = 1
	}
	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum <= 0 {
		limitNum = 10
	}

	var allItems []Budget

	offset := (pageLimit - 1) * limitNum

	err = db.NewSelect().
		Model(&allItems).Limit(limitNum).Relation("Transaction").
		Offset(offset).Order("id ASC").Scan(c.Request.Context())

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"page":  pageLimit,
		"limit": limitNum,
		"items": allItems,
	})
}

func getItemByID(c *gin.Context) {
	itemID := c.Param("id")
	var items []Budget

	if itemID == "" {
		c.JSON(http.StatusNoContent, gin.H{"error": "ID must be present"})
	}

	err := db.NewSelect().Model(&items).Relation("Transaction").Limit(1).Where("id = ?", itemID).Scan(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func updateExistingItem(c *gin.Context) {
	itemID := c.Param("id")
	var updatedItem Budget

	if itemID == "" {
		c.JSON(http.StatusNoContent, gin.H{"error": "ID must be present"})
	}

	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err := db.NewUpdate().Model(&updatedItem).
		Set("Name = ?", updatedItem.Name).
		Set("Amount = ?", updatedItem.Amount).
		Where("id = ?", itemID).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNoContent, "Unable to Update any record")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task Updated Successfully"})
}

func deleteItem(c *gin.Context) {
	itemID := c.Param("id")
	var removedItems []Budget

	if itemID == "" {
		c.JSON(http.StatusNoContent, gin.H{"error": "ID must be present"})
	}

	_, err := db.NewDelete().Model(&removedItems).Where("id = ?", itemID).Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to delete the record")
		return
	}
	c.JSON(http.StatusOK, "Item deleted successfully!!!")
}
