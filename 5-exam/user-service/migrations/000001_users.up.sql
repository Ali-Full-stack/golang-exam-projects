CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(65) NOT NULL,
    email VARCHAR(65) UNIQUE,
    PASSWORD VARCHAR(255) UNIQUE
);
