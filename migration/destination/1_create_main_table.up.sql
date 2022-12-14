CREATE TABLE IF NOT EXISTS destination_product(
   id SERIAL PRIMARY KEY,
   product_name VARCHAR NOT NULL,
   qty INTEGER NOT NULL,
   selling_price NUMERIC NOT NULL,
   promo_price NUMERIC NOT NULL,
   created_at TIMESTAMP,
   updaated_at TIMESTAMP
);