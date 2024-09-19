CREATE TABLE IF NOT EXISTS users(
    id   varchar(255)  PRIMARY KEY,
    name VARCHAR(65) NOT NULL,
    email VARCHAR(65) UNIQUE,
    password_hash VARCHAR(65) NOT NULL
);
