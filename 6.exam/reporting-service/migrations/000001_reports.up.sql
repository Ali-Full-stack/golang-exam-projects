CREATE TABLE IF NOT EXISTS "reports" (
    "id" VARCHAR PRIMARY KEY,
    "income" FLOAT NOT NULL,
    "expense" FLOAT NOT NULL,
    "net_saving" FLOAT NOT NULL,
    "budget_amount" FLOAT NOT NULL,
    "budget_spent" FLOAT NOT NULL,
    "remaining_budget" FLOAT NOT NULL
);