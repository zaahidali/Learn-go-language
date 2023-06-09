package models

import "github.com/uptrace/bun"

// has content, user_id, and post_id.

type Comment struct {
	bun.BaseModel `bun:",alias:comments"`
	ID            uint64    `bun:"id,pk,autoincrement" json:"id"`
	Content       string    `bun:"content,notnull" json:"content"`
	UserID        uint64    `bun:"user_id,notnull" json:"user_id"`
	PostID        uint64    `bun:"post_id,notnull" json:"post_id"`
	SoftDelete
	// User          *User     `bun:"rel:belongs_to,join:user_id=id" json:"user,omitempty"`
	// Post          *Post     `bun:"rel:belongs_to,join:post_id=id" json:"post,omitempty"`
}

// User:
// 	has_many :posts

// Post:
// 	has_many :comments
// 	belongs_to :user

// Comment:
// 	belongs_to :post
// 	belongs_to :user
