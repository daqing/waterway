create table posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id BIGINT REFERENCES users (id),
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create unique index idx_posts_title on posts(title);
create unique index idx_posts_content on posts(content);
