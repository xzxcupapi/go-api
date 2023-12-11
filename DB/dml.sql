-- Sample Customer Data
INSERT INTO customer (name, phoneNumber, address) VALUES
('John Doe', '1234567890', '123 Main St'),
('Jane Doe', '9876543210', '456 Oak St');

-- Sample Product Data
INSERT INTO product (name, price, unit) VALUES
('Product A', 50, 'Unit'),
('Product B', 30, 'Kg');

-- Sample Employee Data
INSERT INTO employee (name, phoneNumber, address) VALUES
('Alice Smith', '1112223333', '789 Elm St'),
('Bob Johnson', '4445556666', '101 Pine St');

-- Sample Transaction Data
INSERT INTO transaction (billDate, entryDate, finishDate, employeeId, customerId, totalBill) VALUES
('2023-01-01', '2023-01-01', '2023-01-02', 1, 1, 150),
('2023-01-02', '2023-01-02', '2023-01-03', 2, 2, 90);

-- Sample Bill Details Data
INSERT INTO billDetails (billId, productId, productPrice, qty) VALUES
(1, 1, 50, 2),
(1, 2, 30, 1),
(2, 2, 30, 3);
