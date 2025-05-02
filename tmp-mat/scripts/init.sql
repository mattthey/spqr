CREATE TABLE orders (
    id SERIAL NOT NULL PRIMARY KEY,
    customer_id INT,
    order_date DATE
);

CREATE TABLE items (
    id SERIAL NOT NULL PRIMARY KEY,
    order_id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR
);
