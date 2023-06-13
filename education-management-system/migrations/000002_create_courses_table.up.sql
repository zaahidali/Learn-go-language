CREATE TABLE IF NOT EXISTS courses (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  	name TEXT NOT NULL,

	created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
	deleted_at TIMESTAMP
);
