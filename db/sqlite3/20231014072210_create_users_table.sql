create table users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  nickname VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  phone VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  avatar VARCHAR(255) NOT NULL,
  role INTEGER NOT NULL,
  encrypted_password VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create unique index idx_users_username on users (username);
create unique index idx_users_password on users (encrypted_password);
