package main

import "github.com/uptrace/bun"

type Budget struct {
	bun.BaseModel `bun:"budgets,alias:b"`
	ID            int64          `bun:"id,autoincrement,pk"`
	Name          string         `bun:"name"`
	Amount        float64        `bun:"amount"`
	Transaction   []*Transaction `bun:"rel:has-many"`
}

type Transaction struct {
	bun.BaseModel `bun:"transactions"`
	ID            int64   `bun:"id,pk"`
	Amount        float64 `bun:"amount"`
	Type          string  `bun:"type"`
	BudgetID      int64   `bun:"budget_id"` // foregin key of the budget
	Budget        *Budget `bun:"rel:belongs-to"`
}

// write migrations to make the budget id a primary key if it is not
// write a migration to link budget with transcation as has_many and belongs_to associations
