package models

import "github.com/uptrace/bun"

type Post struct {
	bun.BaseModel `bun:",alias:posts"`
	ID            uint64    `bun:"id,pk,autoincrement,notnull" json:"id"`
	Title         string    `bun:"title,unique,notnull" json:"title"`
	Content       string    `bun:"content,notnull" json:"content"`
	UserID        uint64    `bun:"user_id,notnull" json:"user_id"`
	Comment []*Comment `bun:"rel:has-many"`
	SoftDelete
	// User          *User     `bun:"rel:belongs_to,join:user_id=id"` // allow this and ignore this when creating new user
}

// User:
// 	has_many :posts

// Post:
// 	has_many :comments
// 	belongs_to :user

// Comment:
// 	belongs_to :post
// 	belongs_to :user
