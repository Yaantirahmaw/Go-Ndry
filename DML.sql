-- Insert data into customer
INSERT INTO customer (id, name, phone, email, address, join_date, active_member, gender)
VALUES ('C001', 'Arjun Aksa Hanendra', '085899888527', 'arjuna@gmail.com', 'embong wungu 12J', '2014-07-01', TRUE, 'M');
INSERT INTO customer (id, name, phone, email, address, join_date, active_member, gender)
VALUES ('C002', 'Candra Dikka Ariesta', '085158668527', 'candraries@gmail.com', 'Ngagel Jaya Selatan 5A', '2016-12-26', TRUE, 'M');
INSERT INTO customer (id, name, phone, email, address, join_date, active_member, gender)
VALUES ('C003', 'Yanti Rahmawati', '0881036336095', 'yaantirahmaw@gmail.com', 'Bratang Gede 1D 36', '2015-11-20', TRUE, 'F');
INSERT INTO customer (id, name, phone, email, address, join_date, active_member, gender)
VALUES ('C004', 'Caca Ria', '081245678910', 'acariangembira@gmail.com', 'Bratang', '2001-11-20', TRUE, 'F');

-- Insert data into service
INSERT INTO service (service_id, name, description, price, duration)
VALUES ('S001', 'Laundry', 'Regular laundry', 10000, '24H');
INSERT INTO service (service_id, name, description, price, duration)
VALUES ('S002', 'Dry Cleaning', 'Dry cleaning only', 20000, '12H');
INSERT INTO service (service_id, name, description, price, duration)
VALUES ('S003', 'Cuci Strika', 'Dry cleaning + strika', 30000, '12H');

-- Insert data into order
INSERT INTO "order" (order_id, customer_id, order_date, total_amount, status)
VALUES ('D001', 'C003', '2024-02-14', 50000, 'Done');
INSERT INTO "order" (order_id, customer_id, order_date, total_amount, status)
VALUES ('D002', 'C001', '2024-05-12', 40000, 'Done');
INSERT INTO "order" (order_id, customer_id, order_date, total_amount, status)
VALUES ('D003', 'C002', '2024-06-27', 60000, 'Done');

-- Insert data into order_detail
INSERT INTO order_detail (order_id, service_id, customer_id, quantity, subtotal)
VALUES ('D001', 'S001', 'C003', '5KG', 50000);
INSERT INTO order_detail (order_id, service_id, customer_id, quantity, subtotal)
VALUES ('D002', 'S002', 'C001', '2KG', 40000);
INSERT INTO order_detail (order_id, service_id, customer_id, quantity, subtotal)
VALUES ('D003', 'S003', 'C002', '2KG', 60000);
