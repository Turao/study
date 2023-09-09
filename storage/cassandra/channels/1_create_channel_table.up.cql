CREATE TABLE IF NOT EXISTS channels.channel (
  id UUID,
  name text,
  tenancy text,
  created_at timestamp,
  deleted_at timestamp,
  PRIMARY KEY((id), tenancy, created_at)
)
WITH CLUSTERING ORDER BY (tenancy DESC, created_at DESC)
AND cdc=true;
