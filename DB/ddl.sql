-- Master Table
CREATE TABLE IF NOT EXISTS customer (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  phoneNumber VARCHAR(20) NOT NULL,
  address TEXT
);

CREATE TABLE IF NOT EXISTS product (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price INT NOT NULL,
  unit VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS employee (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  phoneNumber VARCHAR(20) NOT NULL,
  address TEXT
);

-- Transaction Table
CREATE TABLE IF NOT EXISTS transaction (
  id SERIAL PRIMARY KEY,
  billDate DATE NOT NULL,
  entryDate DATE NOT NULL,
  finishDate DATE NOT NULL,
  employeeId INT REFERENCES employee(id),
  customerId INT REFERENCES customer(id),
  totalBill INT NOT NULL
);

CREATE TABLE IF NOT EXISTS billDetails (
  id SERIAL PRIMARY KEY,
  billId INT REFERENCES transaction(id),
  productId INT REFERENCES product(id),
  productPrice INT NOT NULL,
  qty INT NOT NULL
);
