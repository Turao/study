CREATE TABLE IF NOT EXISTS users(
  _key UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  id UUID NOT NULL,
  version INTEGER NOT NULL,
  email VARCHAR(256) NOT NULL,
  first_name VARCHAR(64),
  last_name VARCHAR(64),
  tenancy VARCHAR(32),
  created_at TIMESTAMP,
  deleted_at TIMESTAMP
);

ALTER TABLE users ADD CONSTRAINT users_unique_id_version UNIQUE (id, version);

ALTER TABLE users REPLICA IDENTITY FULL;
