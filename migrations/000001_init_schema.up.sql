CREATE TABLE IF NOT EXISTS events(
  id uuid primary key,
  title varchar(255),
  description text,
  start_at TIMESTAMP,
  finish_at TIMESTAMP,
  user_id uuid
);

