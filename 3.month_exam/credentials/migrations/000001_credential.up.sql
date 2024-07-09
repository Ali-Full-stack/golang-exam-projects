CREATE TABLE "clients" (
  "id" uuid PRIMARY KEY,
  "name" varchar(65),
  "email" varchar(65) UNIQUE,
  "phone" varchar(65),
  "created_at" varchar(255)
);

CREATE TABLE "client_location" (
  "client_id" uuid,
  "city" varchar(65),
  "region" varchar(65),
  "home_address" varchar(65)
);

CREATE TABLE "drivers" (
  "id" uuid PRIMARY KEY,
  "name" varchar(65),
  "email" varchar(65) UNIQUE,
  "phone" varchar(65),
  "working_region" varchar(65),
  "vehicle" varchar(65),
  "status" varchar(65),
  "hired_at" varchar(65)
);

CREATE TABLE "driver_location" (
  "driver_id" uuid,
  "city" varchar(65),
  "region" varchar(65),
  "home_address" varchar(65)
);

CREATE TABLE "admins" (
  "id" uuid PRIMARY KEY,
  "name" varchar(65),
  "email" varchar(65) UNIQUE,
  "password" varchar(255),
  "role" varchar(65)
);

ALTER TABLE "client_location" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");

ALTER TABLE "driver_location" ADD FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id");