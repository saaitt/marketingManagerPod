CREATE TABLE products
(
    id        SERIAL PRIMARY KEY,
    title     varchar(200),
    page_link varchar(1000),
    user_id   int,
    FOREIGN KEY (user_id) REFERENCES users (id)

);

CREATE TABLE users
(
    id        SERIAL PRIMARY KEY,
    username  varchar(50),
    password  varchar(50),
    user_type varchar(20)
);

CREATE TABLE marketing_products
(
    id         SERIAL PRIMARY KEY,
    UsageCount INT,
    uuid       varchar(200) NOT NULL,
    productId  INT,
    userId     INT,
    FOREIGN KEY (productId) REFERENCES products (id),
    FOREIGN KEY (userId) REFERENCES users (id)
);
