CREATE TABLE "client_account" (
  "client_id" varchar(255) PRIMARY KEY,
  "card_number" varchar(16),
  "balance" float   not null check (balance > 0)
);

CREATE TABLE "driver_account" (
  "driver_id" varchar(255) PRIMARY KEY,
  "card_number" varchar(16),
  "balance" float not null check (balance > 0)
);

CREATE TABLE "main_account" (
  "id" varchar(255) PRIMARY KEY,
  "account_number" varchar(65),
  "balance" float   not null check (balance > 0)
);
