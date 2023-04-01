CREATE TABLE IF NOT EXISTS messages.message (
  id UUID,
  content text,
  tenancy text,
  created_at timestamp,
  deleted_at timestamp,
  PRIMARY KEY(id)
) WITH cdc=true;
