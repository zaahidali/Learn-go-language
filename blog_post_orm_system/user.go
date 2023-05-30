package main

import (
	"time"

	"github.com/uptrace/bun"
)

// User:
// 	has_many :posts

// Post:
// 	has_many :comments
// 	belongs_to :user

// Comment:
// 	belongs_to :post
// 	belongs_to :user

// has name and email

type User struct {
	bun.BaseModel `bun:",alias:users"`
	ID            uint64    `bun:"id,pk,autoincrement,notnull"`
	Name          string    `bun:"name,notnull"`
	Email         string    `bun:"email,notnull,unique"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
	DeletedAt     time.Time `bun:"deleted_at"`
	Post          []*Post   `bun:"rel:has-many"`
}
