CREATE TABLE IF NOT EXISTS transactions (
  id uuid PRIMARY KEY,
  account_id uuid,
  operation_type_id integer,
  amount numeric,
  event_date timestamp,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp,

  CONSTRAINT fk_transactions_accounts FOREIGN KEY(account_id) REFERENCES accounts(id),
  CONSTRAINT fk_transactions_transaction_types FOREIGN KEY(operation_type_id) REFERENCES transaction_types(id)
);

