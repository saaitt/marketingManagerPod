CREATE TABLE marketing_products (
    id SERIAL PRIMARY KEY,
    countUsage INT,
    uuid varchar(200) NOT NULL,
    productId INT ,
    FOREIGN KEY (productId) REFERENCES products(id)
);