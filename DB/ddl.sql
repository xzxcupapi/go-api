-- Tabel Master Customer
CREATE TABLE customers (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  phoneNumber VARCHAR(20) NOT NULL,
  address VARCHAR(255) NOT NULL
);

-- Tabel Master Product
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price INT NOT NULL,
  unit VARCHAR(20) NOT NULL
);

-- Tabel Master Employee
CREATE TABLE employees (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  phoneNumber VARCHAR(20) NOT NULL,
  address TEXT NOT NULL
);

-- Tabel Transaksi
CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  billDate DATE NOT NULL,
  entryDate DATE NOT NULL,
  finishDate DATE NOT NULL,
  employeeId INT REFERENCES employees(id),
  customerId INT REFERENCES customers(id)
);

-- Tabel Detail Transaksi
CREATE TABLE billDetails (
  id SERIAL PRIMARY KEY,
  billId INT REFERENCES transactions(id),
  productId INT REFERENCES products(id),
  productPrice INT NOT NULL,
  qty INT NOT NULL
);
