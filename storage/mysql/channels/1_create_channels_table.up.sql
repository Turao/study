CREATE TABLE IF NOT EXISTS channels(
  id VARCHAR(36) PRIMARY KEY,
  name VARCHAR(64),
  tenancy VARCHAR(32),
  created_at TIMESTAMP,
  deleted_at TIMESTAMP
);
