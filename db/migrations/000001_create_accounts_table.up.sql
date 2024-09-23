CREATE TABLE IF NOT EXISTS accounts (
  id uuid PRIMARY KEY,
  document_number VARCHAR UNIQUE
);
