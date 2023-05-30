package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var db *bun.DB

func main() {
	fmt.Println("********************* Budget Management System *********************")

	/*
		POST /budgets - Create a new budget. This should accept a JSON body with the budget's name and amount. (DONE)
		GET /budgets - Get a list of all budgets. (Done)
		GET /budgets/:id - Get the details of a specific budget. (Done)
		PUT /budgets/:id - Update a specific budget. This should accept a JSON body with the new details of the budget.
		DELETE /budgets/:id - Delete a specific budget.

		POST /transactions - Record a new transaction. This should accept a JSON body with the transaction's details including amount, type (income or expense), and associated budget id.
		GET /transactions - Get a list of all transactions.
		GET /transactions/:id - Get the details of a specific transaction.
		PUT /transactions/:id - Update a specific transaction. This should accept a JSON body with the new details of the transaction.
		DELETE /transactions/:id - Delete a specific transaction.

	*/

	// 1. Make a database connection
	ctx := context.Background()
	dsn := "postgres://postgres:postgres@localhost:5432/budget_management_system?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return
	}
	fmt.Println("Connected to the database", ctx)

	// 2. Create Model/Table
	// err = db.ResetModel(ctx, (*Budget)(nil))
	// _, err = db.NewDropTable().Model((*Budget)(nil)).IfExists().Exec(ctx)
	// if err != nil {
	// 	fmt.Println("Unable to Recreate the Table")
	// 	return
	// }
	fmt.Println("Budget Table created successfully!!!")

	// 3. Add router and register routes
	router := gin.New()
	router.GET("/", welcomePage)
	router.POST("/budgets", addNewItem)
	router.GET("/budgets", getPaginatedItems)
	router.GET("/budgets/:id", getItemByID)
	router.PUT("/budgets/:id", updateExistingItem)
	router.DELETE("/budgets/:id", deleteItem)

	router.GET("/transactions", getTransactions)
	router.GET("/transactions/:id", getTransactionByID)
	router.POST("/transactions", addNewTransaction)
	router.DELETE("/transactions/:id", deleteTransaction)
	router.Run()
}
