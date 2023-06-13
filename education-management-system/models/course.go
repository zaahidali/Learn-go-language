package models

import (
	"github.com/uptrace/bun"
)

type Course struct {
	bun.BaseModel `bun:",alias:courses"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
	SoftDelete
}
