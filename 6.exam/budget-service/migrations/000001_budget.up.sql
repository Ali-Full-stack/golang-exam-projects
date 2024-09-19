CREATE TABLE IF NOT EXISTS "budgets" (
    "id"  VARCHAR  PRIMARY KEY,
    "user_id" VARCHAR NOT NULL,
    "category" VARCHAR NOT NULL,
    "amount" FLOAT NOT NULL,
    "spent" FLOAT NOT NULL,
    "currency" VARCHAR NOT NULL
);