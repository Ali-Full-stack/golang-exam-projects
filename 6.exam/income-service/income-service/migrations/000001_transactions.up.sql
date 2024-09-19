CREATE TABLE IF NOT EXISTS  "transactions" (
  "id" varchar PRIMARY KEY ,
  "user_id" varchar NOT NULL,
  "type" varchar NOT NULL,
  "category" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "amount" FLOAT NOT NULL,
  "date" varchar NOT NULL
);