package models

import "github.com/uptrace/bun"

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
	Post          []*Post   `bun:"rel:has-many"`
	SoftDelete
}
