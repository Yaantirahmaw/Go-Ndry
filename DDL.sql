CREATE DATABASE go_ndry;

CREATE TABLE customer (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    join_date DATE NOT NULL,
    active_member BOOLEAN NOT NULL,
    gender CHAR(1) NOT NULL
);

CREATE TABLE service (
    service_id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    duration VARCHAR(5) NOT NULL
);

CREATE TABLE "order" (
    order_id VARCHAR(10) PRIMARY KEY,
    customer_id VARCHAR(10),
    order_date DATE NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customer(id)
);

CREATE SEQUENCE order_detail_id_seq;

CREATE TABLE order_detail (
    detail_id INT PRIMARY KEY DEFAULT nextval('order_detail_id_seq'),
    order_id VARCHAR(10),
    service_id VARCHAR(10),
    customer_id VARCHAR(10),
    quantity VARCHAR(5) NOT NULL,
    subtotal DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES "order"(order_id),
    FOREIGN KEY (service_id) REFERENCES service(service_id),
    FOREIGN KEY (customer_id) REFERENCES customer(id)
);
