package main

import (
	"time"

	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:",alias:posts"`
	ID            uint64    `bun:"id,pk,autoincrement,notnull" json:"id"`
	Title         string    `bun:"title,unique,notnull" json:"title"`
	Content       string    `bun:"content,notnull" json:"content"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
	DeletedAt     time.Time `bun:"deleted_at" json:"deleted_at"`
	UserID        uint64    `bun:"user_id,notnull" json:"user_id"`
	// User          *User     `bun:"rel:belongs_to,join:user_id=id"` // allow this and ignore this when creating new user
	Comment []*Comment `bun:"rel:has-many"`
}

// User:
// 	has_many :posts

// Post:
// 	has_many :comments
// 	belongs_to :user

// Comment:
// 	belongs_to :post
// 	belongs_to :user
