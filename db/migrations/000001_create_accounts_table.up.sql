CREATE TABLE IF NOT EXISTS accounts (
  id uuid PRIMARY KEY,
  document_number VARCHAR UNIQUE,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);
