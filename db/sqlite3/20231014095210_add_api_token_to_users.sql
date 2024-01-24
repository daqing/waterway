ALTER TABLE users ADD api_token VARCHAR(255);
create unique index idx_api_token on users (api_token);
