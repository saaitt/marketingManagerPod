CREATE TABLE marketing_products (
    id SERIAL PRIMARY KEY,
    countUsage INT,
    UUID string NOT NULL,
    productId SERIAL FOREIGN KEY REFERENCES products(id),

);