package models

import (
	"github.com/uptrace/bun"
)

type Student struct {
	bun.BaseModel `bun:"alias:students"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull" json:"name"`
	SoftDelete
}
