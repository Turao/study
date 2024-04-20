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

CREATE TABLE IF NOT EXISTS group_member(
  _key UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  group_id UUID NOT NULL,
  group_version INTEGER NOT NULL,
  member_id UUID NOT NULL
);

ALTER TABLE group_member ADD CONSTRAINT group_member_unique_group_id_group_version_member_id UNIQUE (group_id, group_version, member_id);

ALTER TABLE group_member REPLICA IDENTITY FULL;