CREATE TABLE IF NOT EXISTS memberships(
  _key VARCHAR(36) DEFAULT (uuid()) PRIMARY KEY,
  channel_id VARCHAR(36) NOT NULL,
  user_id VARCHAR(36) NOT NULL,
  version INTEGER NOT NULL,
  tenancy VARCHAR(32),
  created_at TIMESTAMP,
  deleted_at TIMESTAMP
);

ALTER TABLE memberships ADD CONSTRAINT unique_channel_id_user_id_version UNIQUE (channel_id, user_id, version);