create table posts (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES users (id),
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create unique index idx_posts_title on posts(title);
create unique index idx_posts_content on posts(content);
