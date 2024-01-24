create table tag_relation (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  tag_id BIGINT REFERENCES tags(id),
  relation_type VARCHAR(255) NOT NULL,
  relation_id BIGINT  NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)

-- create index idx_tag_relation on tag_relation(tag_id, relation_type, relation_id)
