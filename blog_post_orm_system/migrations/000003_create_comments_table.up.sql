CREATE TABLE comments (
  id BIGSERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  user_id BIGINT NOT NULL REFERENCES users(id),
  post_id BIGINT NOT NULL REFERENCES posts(id),
  created_at TIMESTAMP DEFAULT current_timestamp,
  updated_at TIMESTAMP DEFAULT current_timestamp,
  deleted_at TIMESTAMP
);

-- // User:
-- // 	has_many :posts

-- // Post:
-- // 	has_many :comments
-- // 	belongs_to :user

-- // Comment:
-- // 	belongs_to :post
-- // 	belongs_to :user
