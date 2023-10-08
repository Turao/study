CREATE TABLE IF NOT EXISTS channels(
  _key VARCHAR(36) DEFAULT (uuid()) PRIMARY KEY,
  id VARCHAR(36),
  version INTEGER,
  name VARCHAR(64),
  tenancy VARCHAR(32),
  created_at TIMESTAMP,
  deleted_at TIMESTAMP
);
