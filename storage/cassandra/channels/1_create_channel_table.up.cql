CREATE TABLE IF NOT EXISTS channels.channel (
  id UUID,
  name text,
  tenancy text,
  created_at timestamp,
  deleted_at timestamp,
  PRIMARY KEY(id)
) WITH cdc=true;
