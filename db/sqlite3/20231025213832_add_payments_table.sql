create table payments (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid VARCHAR(255) NOT NULL,
  user_id BIGINT NOT NULL references users(id),
  goods_type VARCHAR(255) NOT NULL,
  goods_id BIGINT NOT NULL,
  cent INT NOT NULL,
  action VARCHAR(255) NOT NULL,
  note VARCHAR(255) NOT NULL DEFAULT '',
  status VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create unique index idx_payments_uuid on payments(uuid);
