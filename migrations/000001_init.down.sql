ALTER TABLE IF EXISTS "product_reservation" DROP CONSTRAINT IF EXISTS "product_reservation_unique_code_id_fkey";
ALTER TABLE IF EXISTS "product_reservation" DROP CONSTRAINT IF EXISTS "product_reservation_warehouse_id_fkey";
ALTER TABLE IF EXISTS "inventory" DROP CONSTRAINT IF EXISTS "inventory_option_id_fkey";
ALTER TABLE IF EXISTS "inventory" DROP CONSTRAINT IF EXISTS "inventory_warehouse_id_fkey";
ALTER TABLE IF EXISTS "option" DROP CONSTRAINT IF EXISTS "option_product_id_fkey";
ALTER TABLE IF EXISTS "option" DROP CONSTRAINT IF EXISTS "option_size_fkey";

DROP TABLE IF EXISTS "product_reservation";
DROP TABLE IF EXISTS "inventory";
DROP TABLE IF EXISTS "option";
DROP TABLE IF EXISTS "size";
DROP TABLE IF EXISTS "warehouse";
DROP TABLE IF EXISTS "product";
