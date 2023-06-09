CREATE TABLE IF NOT EXISTS enrollments (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name TEXT NOT NULL,
	student_id BIGINT REFERENCES Students(id) ON DELETE CASCADE,
	course_id BIGINT REFERENCES Courses(id) ON DELETE CASCADE,

	created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
	deleted_at TIMESTAMP
);
