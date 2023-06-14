CREATE TABLE IF NOT EXISTS posts (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL UNIQUE,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  deleted_at TIMESTAMP,
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE
);

-- // User:
-- // 	has_many :posts

-- // Post:
-- // 	has_many :comments
-- // 	belongs_to :user

-- // Comment:
-- // 	belongs_to :post
-- // 	belongs_to :user
