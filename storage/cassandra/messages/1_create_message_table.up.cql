CREATE TABLE IF NOT EXISTS messages.message (
  id UUID,
  author UUID,
  channel text,
  content text,
  tenancy text,
  created_at timestamp,
  deleted_at timestamp,
  PRIMARY KEY(id, author, channel)
) WITH cdc=true;
