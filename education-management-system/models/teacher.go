package models

type Teacher struct {
	ID       int64  `bun:"id,pk,autoincrement"`
	Name     string `bun:"name,notnull" json:"name"`
	CourseID int64  `bun:"rel:belongs-to,join:course_id=id"`
	SoftDelete
}
