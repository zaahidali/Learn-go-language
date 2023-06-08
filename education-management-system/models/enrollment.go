package models

type Enrollment struct {
	ID        int64 `bun:"id,pk,autoincrement"`
	StudentID int64 `bun:"student_id" json:"student_id"`
	CourseID  int64 `bun:"course_id" json:"course_id"`
	SoftDelete

	Student *Student `bun:"rel:belongs-to,join:student_id=id"`
	Course  *Course  `bun:"rel:belongs-to,join:course_id=id"`
}
