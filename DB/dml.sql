-- Sample Data Customer
INSERT INTO customers (id, name, phoneNumber, address) VALUES
('1', 'John Doe', '081234567890', 'Jl. Mawar No. 123'),
('2', 'Jane Doe', '082345678901', 'Jl. Melati No. 456');

-- Sample Data Product
    INSERT INTO products (id, name, quantity, unit, price) VALUES
    ('1', 'Laundry Reguler', 1, 'KG', 15000),
	('2', 'Laundry Express', 1, 'KG', 20000),;

-- Sample Data Employee
INSERT INTO employees (id, name, phoneNumber, address) VALUES
('1', 'Alice', '083456789012', 'Jl. Dahlia No. 789'),
('2', 'Bob', '084567890123', 'Jl. Anggrek No. 012');

-- Sample Data Transaction
INSERT INTO transactions (billDate, entryDate, finishDate, employeeId, customerId) VALUES
('2023-01-01', '2023-01-02', '2023-01-03', 1, 1),
('2023-01-02', '2023-01-03', '2023-01-04', 2, 2);

-- Sample Data Bill Details
INSERT INTO billDetails (billId, productId, productPrice, qty) VALUES
(1, 1, 150000, 2),
(2, 2, 200000, 1);
