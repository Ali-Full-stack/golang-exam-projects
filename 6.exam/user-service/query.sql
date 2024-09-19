-- name: CreateUser :one
INSERT INTO users (id,name, email, password_hash) 
VALUES ($1, $2, $3, $4)
RETURNING id, name, email;

-- name: LoginUser :one
SELECT password_hash FROM users
WHERE id = $1;

-- name: ForgotPassword :one
SELECT  email  FROM users
WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE users
SET password_hash = $1
WHERE email = $2 ;

-- name: LogOutUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserByID :one
SELECT id, name , email FROM users
WHERE id = $1 ;
