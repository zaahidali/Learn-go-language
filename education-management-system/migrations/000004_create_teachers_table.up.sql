CREATE TABLE IF NOT EXISTS teachers (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    course_id BIGINT REFERENCES courses(id) ON DELETE CASCADE,

    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);
