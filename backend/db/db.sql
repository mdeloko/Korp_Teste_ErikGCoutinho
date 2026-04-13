CREATE TYPE status AS ENUM ('opened','closed');

CREATE TABLE invoices (
    id SERIAL PRIMARY KEY,
    invoice_status status NOT NULL DEFAULT 'opened'
);

CREATE TABLE products(
    code VARCHAR(10) PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    amount INTEGER NOT NULL
);

CREATE TABLE products_to_invoices(
    id BIGSERIAL PRIMARY KEY,
    invoice_id INTEGER NOT NULL REFERENCES invoices(id),
    product_id VARCHAR(10) NOT NULL REFERENCES products(code),
    amount INTEGER NOT NULL DEFAULT 1
);

INSERT INTO invoices (invoice_status) VALUES ('closed');
INSERT INTO invoices (invoice_status) VALUES ('opened');

INSERT INTO products (code,description,amount) VALUES ('001', 'Banana', 4);
INSERT INTO products (code,description,amount) VALUES ('002', 'Pera', 5);
