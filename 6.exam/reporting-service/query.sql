-- name: CreateReport :exec
INSERT INTO reports (id, income, expense, net_saving, budget_amount, budget_spent, remaining_budget)
VALUES ($1, $2, $3, $4, $5, $6, $7);
