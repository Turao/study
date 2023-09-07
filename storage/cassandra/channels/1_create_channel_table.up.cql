CREATE TABLE IF NOT EXISTS channels.channel (
  id UUID,
  name text,
  tenancy text,
  created_at timestamp,
  deleted_at timestamp,
  PRIMARY KEY((tenancy))
)
WITH CLUSTERING ORDER BY (created_at DESC)
AND cdc=true;
