CREATE TABLE IF NOT EXISTS hotels(
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(65) UNIQUE,
    rating DECIMAL(2,1) NOT NULL,
    city VARCHAR(65) NOT NULL,
    region VARCHAR(65) NOT NULL,
    street VARCHAR(65) NOT NULL
);
CREATE TABLE IF NOT EXISTS rooms (
    hotel_id VARCHAR(255) NOT NULL ,
    type VARCHAR(65) NOT NULL,
    pricePerNight DECIMAL(7,2) NOT NULL,
    totalRooms INTEGER NOT NULL CHECK (totalRooms >= 0),
    FOREIGN KEY(hotel_id) REFERENCES hotels(id) ON DELETE CASCADE
)
 
