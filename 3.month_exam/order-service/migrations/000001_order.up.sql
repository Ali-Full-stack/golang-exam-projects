CREATE TABLE "products" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255),
  "category" varchar(65),
  "quantity" int check(quantity > 0),
  "price" float check(price > 0),
  "created_at" varchar(100),
  "expired_at" varchar(100)
);
