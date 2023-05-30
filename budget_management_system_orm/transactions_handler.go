package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*

	POST /transactions - Record a new transaction. This should accept a JSON body with the transaction's details including amount, type (income or expense), and associated budget id.
	GET /transactions - Get a list of all transactions.
	GET /transactions/:id - Get the details of a specific transaction.
	PUT /transactions/:id - Update a specific transaction. This should accept a JSON body with the new details of the transaction.
	DELETE /transactions/:id - Delete a specific transaction.

*/

// func getTransactions(c *gin.Context) {
// 	var transactions []Transaction
// 	_, err := db.NewSelect().Model(&transactions).Relation("Budget").Order("id ASC").Exec(c.Request.Context())
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, transactions)
// }

func getTransactions(c *gin.Context) {
	var transactions []Transaction
	err := db.NewSelect().Model(&transactions).Relation("Budget").Order("id ASC").Scan(c.Request.Context())
	if err != nil {
		log.Printf("Error getting transactions: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Transactions: %+v", transactions)
	c.JSON(http.StatusOK, transactions)
}

func addNewTransaction(c *gin.Context) {
	/*
	   POST /transactions - Record a new transaction. This should accept a JSON body with the transaction's details including amount, type (income or expense), and associated budget id.
	*/
	var newTransaction Transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err := db.NewInsert().Model(&newTransaction).Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Record succesfully created!!!", "result": newTransaction})
}

func getTransactionByID(c *gin.Context) {
	transactionIDStr := c.Param("id")
	transactionID, err := strconv.Atoi(transactionIDStr) // Converted string to integer

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	var transaction Transaction

	err = db.NewSelect().Model(&transaction).Relation("Budget").Where("transaction.id = ?", transactionID).Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)

}

func deleteTransaction(c *gin.Context) {
	transactionIDStr := c.Param("id")
	transactionID, err := strconv.Atoi(transactionIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Transaction ID"})
		return
	}

	var transaction Transaction
	_, err = db.NewDelete().Model(&transaction).Where("id = ?", transactionID).Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record successfully deleted!!!", "result": transaction})
}
