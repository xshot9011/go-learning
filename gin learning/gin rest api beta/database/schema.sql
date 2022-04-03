CREATE TABLE Account(id SERIAL UNIQUE, email VARCHAR(255) NOT NULL, password_hash VARCHAR(255), is_active BOOLEAN, created TIMESTAMP, updated TIMESTAMP);

CREATE TABLE Item(id SERIAL UNIQUE, title VARCHAR(255) NOT NULL, notes TEXT, seller_id INTEGER, price_in_cents INTEGER, FOREIGN KEY (seller_id) REFERENCES Account(id) ON DELETE CASCADE, created TIMESTAMP, updated TIMESTAMP);

CREATE TABLE Purchase(id SERIAL UNIQUE,customer_id INTEGER,item_id INTEGER,price_in_cents INTEGER,title VARCHAR(255) NOT NULL,FOREIGN KEY (customer_id) REFERENCES Account (id),FOREIGN KEY (item_id) REFERENCES Item (id),created TIMESTAMP,updated TIMESTAMP);