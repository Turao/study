CREATE TABLE IF NOT EXISTS messages.message (
  id UUID,
  version int,
  author UUID,
  channel text,
  content text,
  tenancy text,
  created_at timestamp,
  deleted_at timestamp,
  PRIMARY KEY((channel), created_at)
)
WITH CLUSTERING ORDER BY (created_at DESC)
AND cdc=true;