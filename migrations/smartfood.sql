CREATE TABLE "users" (
  "id" VARCHAR NOT NULL,
  "name" text,
  "created_at" timestamp
);
CREATE TABLE "foods" (
  "id" SERIAL PRIMARY KEY,
  "name" text,
  "description" text,
  "category" text,
  "image" text
);
CREATE TABLE "detection_history" (
  "id" SERIAL PRIMARY KEY,
  "result" jsonb,
  "created_at" timestamptz,
  "user_id" text
);
CREATE TABLE "shops" (
  "id" SERIAL PRIMARY KEY,
  "name" text,
  "location" text,
  "gmaps_link" text,
  "latitude" float,
  "longitude" float,
  "image" text
);
CREATE TABLE "shop_foods" (
  "id" SERIAL PRIMARY KEY,
  "shop_id" integer,
  "food_id" integer,
  "price" float
);

ALTER TABLE "shop_foods"
ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");
ALTER TABLE "shop_foods"
ADD FOREIGN KEY ("food_id") REFERENCES "foods" ("id");
ALTER TABLE "detection_history"
ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
