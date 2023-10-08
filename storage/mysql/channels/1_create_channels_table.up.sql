CREATE TABLE IF NOT EXISTS channels(
  _key VARCHAR(36) DEFAULT (uuid()) PRIMARY KEY,
  id VARCHAR(36) NOT NULL,
  version INTEGER NOT NULL,
  name VARCHAR(64),
  tenancy VARCHAR(32),
  created_at TIMESTAMP,
  deleted_at TIMESTAMP
);

ALTER TABLE channels ADD CONSTRAINT unique_id_version UNIQUE (id, version);