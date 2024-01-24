create table media_files (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id BIGINT NOT NULL references users(id),
  filename VARCHAR(255) NOT NULL,
  mime VARCHAR(255) NOT NULL,
  bytes BIGINT NOT NULL,
  expired_at TIMESTAMP NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create unique index idx_medie_files_filename on media_files(filename);
