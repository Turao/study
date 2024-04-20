CREATE TABLE IF NOT EXISTS groups(
  _key UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  id UUID NOT NULL,
  version INTEGER NOT NULL,
  name VARCHAR(64),
  tenancy VARCHAR(32),
  created_at TIMESTAMP,
  deleted_at TIMESTAMP
);

ALTER TABLE groups ADD CONSTRAINT groups_unique_id_version UNIQUE (id, version);

ALTER TABLE groups REPLICA IDENTITY FULL;
