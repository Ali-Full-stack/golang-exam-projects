-- name: CreateBudget :one
INSERT INTO budgets(id, user_id, category, amount, spent, currency)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING id ;

-- name: UpdateBudgetByID :exec
UPDATE budgets
SET  amount = $1
WHERE id = $2 and user_id = $3 ;

-- name: GetBudgets :many
SELECT id, category, amount, spent, currency FROM budgets
WHERE user_id = $1 ;

-- name: DeleteBudgetByID :exec
DELETE FROM budgets
WHERE id = $1 and user_id = $2;

-- name: GetSpentFromBudget :one
SELECT spent from budgets
WHERE id = $1 and user_id = $2;

-- name: GetBudgetReport :one
SELECT SUM(amount) as total_budget, SUM(spent) as total_spent  FROM budgets
WHERE user_id = $1;

-- name: AddSpentAmount :exec
UPDATE budgets
SET spent = spent + $1
WHERE id = $2 and amount >= spent + $1 ;


