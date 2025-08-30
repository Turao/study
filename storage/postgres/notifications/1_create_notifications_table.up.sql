CREATE TABLE IF NOT EXISTS notifications(
  _key UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  id UUID NOT NULL,
  type VARCHAR(256) NOT NULL,
  recipient VARCHAR(256),
  subject VARCHAR(256),
  content VARCHAR(1024),
  created_at TIMESTAMP,
  external_reference_id VARCHAR(256)
);

ALTER TABLE notifications ADD CONSTRAINT notifications_unique_id UNIQUE (id);

ALTER TABLE notifications REPLICA IDENTITY FULL;
