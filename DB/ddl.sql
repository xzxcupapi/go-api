-- Tabel Produk
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    unit VARCHAR(20) NOT NULL
);

-- Tabel Karyawan
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phoneNumber VARCHAR(20) NOT NULL,
    address TEXT NOT NULL
);

-- Tabel Transaksi
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    billDate DATE NOT NULL,
    entryDate DATE NOT NULL,
    finishDate DATE NOT NULL,
    employeeID INT REFERENCES employees(id) NOT NULL,
    customerID INT REFERENCES customers(id) NOT NULL,
    totalBill DECIMAL(10, 2) NOT NULL
);

-- Tabel DetailTransaksi
CREATE TABLE IF NOT EXISTS transactionDetails (
    id SERIAL PRIMARY KEY,
    billID INT REFERENCES transactions(id) NOT NULL,
    productID INT REFERENCES products(id) NOT NULL,
    productName VARCHAR(255) NOT NULL,
    productPrice INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0)
);
