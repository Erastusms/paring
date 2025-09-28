-- User Table
CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    passwordhash TEXT NOT NULL,
    createdat TIMESTAMP DEFAULT NOW()
);

-- Product Table
CREATE TABLE products (
    productid SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    price NUMERIC(12,2) NOT NULL,
    stockquantity INT DEFAULT 0,
    createdat TIMESTAMP DEFAULT NOW()
);

-- Order Table
CREATE TABLE orders (
    orderid SERIAL PRIMARY KEY,
    userid INT NOT NULL REFERENCES users(userid),
    orderdate TIMESTAMP DEFAULT NOW(),
    totalamount NUMERIC(12,2) NOT NULL,
    status VARCHAR(50) DEFAULT 'PENDING'
);

-- OrderItem Table
CREATE TABLE orderitems (
    orderitemid SERIAL PRIMARY KEY,
    orderid INT NOT NULL REFERENCES orders(orderid) ON DELETE CASCADE,
    productid INT NOT NULL REFERENCES products(productid),
    quantity INT NOT NULL,
    price NUMERIC(12,2) NOT NULL
);
