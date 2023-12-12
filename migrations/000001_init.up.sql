CREATE TABLE "product" (
  "id" uuid PRIMARY KEY,
  "name" TEXT NOT NULL
);

CREATE TABLE "size" (
  "name" TEXT UNIQUE PRIMARY KEY
);

CREATE TABLE "option" (
  "product_id" uuid,
  "size" TEXT,
  "id" TEXT UNIQUE PRIMARY KEY NOT NULL
);

CREATE TABLE "product_reservation" (
  "id" uuid PRIMARY KEY,
  "unique_code_id" TEXT,
  "warehouse_id" uuid,
  "quantity_reserved" INT NOT NULL,
  "reservation_date" DATE
);

CREATE TABLE "warehouse" (
  "id" uuid PRIMARY KEY,
  "office_name" TEXT
);

CREATE TABLE "inventory" (
  "option_id" TEXT,
  "warehouse_id" uuid,
  "quantity" int
);

ALTER TABLE "option" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "option" ADD FOREIGN KEY ("size") REFERENCES "size" ("name");

ALTER TABLE "inventory" ADD FOREIGN KEY ("option_id") REFERENCES "option" ("id");

ALTER TABLE "inventory" ADD FOREIGN KEY ("warehouse_id") REFERENCES "warehouse" ("id");

ALTER TABLE "product_reservation" ADD FOREIGN KEY ("warehouse_id") REFERENCES "warehouse" ("id");

ALTER TABLE "product_reservation" ADD FOREIGN KEY ("unique_code_id") REFERENCES "option" ("id");
