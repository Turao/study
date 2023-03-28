CREATE TABLE IF NOT EXISTS users(
  id UUID PRIMARY KEY,
  email VARCHAR(256) NOT NULL,
  first_name VARCHAR(64),
  last_name VARCHAR(64),
  tenancy VARCHAR(32),
  created_at TIMESTAMP,
  deleted_at TIMESTAMP
);

ALTER TABLE users REPLICA IDENTITY FULL;
