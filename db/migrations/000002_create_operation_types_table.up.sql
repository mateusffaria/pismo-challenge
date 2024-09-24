CREATE TABLE IF NOT EXISTS transaction_types (
  id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  description VARCHAR,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

INSERT INTO transaction_types (description, created_at, updated_at)
  VALUES ('PURCHASE', NOW(), NOW());
INSERT INTO transaction_types (description, created_at, updated_at)
  VALUES ('INSTALLMENT_PURCHASE', NOW(), NOW());
INSERT INTO transaction_types (description, created_at, updated_at)
  VALUES ('WITHDRAWAL', NOW(), NOW());
INSERT INTO transaction_types (description, created_at, updated_at)
  VALUES ('PAYMENT', NOW(), NOW());


