-- name: CreateTransaction :one
INSERT INTO transactions (id,user_id,type, category, currency ,amount, date) 
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id ;

-- name: UpdateTransactionByID :exec
UPDATE transactions
SET type = $2, category = $3, currency = $4, amount = $5, date = $6
WHERE id = $1;

-- name: DeleteTransactionByID :exec
DELETE FROM transactions
WHERE id = $1;

-- name: GetTransactionByID :one
SELECT * FROM transactions
WHERE id = $1;

-- name: GetTransactionsByCategory :many
SELECT * FROM transactions
WHERE category = $1 AND user_id = $2;

-- name: GetTransactionByDate :many
SELECT * FROM transactions
WHERE type = $1
AND user_id = $2 and date BETWEEN $3 AND $4 ;

-- name: GetIncomes :one
SELECT SUM(amount)  FROM transactions
WHERE type =$1 AND user_id = $2;

-- name: GetExpenses :one
SELECT SUM(amount) FROM transactions
WHERE type =$1 AND user_id = $2;

-- name: GetReportByType :many
SELECT category, SUM(amount) as total_aount
FROM transactions
WHERE type = $1 AND user_id = $2 GROUP BY category;

-- name: GetReportByCategory :one
SELECT category, amount 
FROM transactions
WHERE type = $1 AND category = $2 AND user_id - $3;
