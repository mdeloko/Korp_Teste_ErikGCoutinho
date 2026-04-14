CREATE TYPE status AS ENUM ('opened','closed');

CREATE TABLE invoices (
    id SERIAL PRIMARY KEY,
    invoice_status status NOT NULL DEFAULT 'opened'
);

CREATE TABLE products(
    code VARCHAR(10) PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    amount INTEGER NOT NULL CHECK (amount >=0),
    is_active BOOLEAN DEFAULT TRUE NOT NULL
);

CREATE TABLE products_to_invoices(
    id BIGSERIAL PRIMARY KEY,
    invoice_id INTEGER NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
    product_id VARCHAR(10) NOT NULL REFERENCES products(code),
    amount INTEGER NOT NULL DEFAULT 1
);

CREATE OR REPLACE FUNCTION restore_product_amount()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE products
    SET amount = amount + OLD.amount
    WHERE code = OLD.product_id;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_restore_stock_on_delete
AFTER DELETE ON products_to_invoices
FOR EACH ROW EXECUTE FUNCTION restore_product_amount();

CREATE OR REPLACE FUNCTION decrease_product_amount()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE products
    SET amount = amount - NEW.amount
    WHERE code = NEW.product_id AND amount >= NEW.amount;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Insufficient stock for product %', NEW.product_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_decrease_stock_on_insert
AFTER INSERT ON products_to_invoices
FOR EACH ROW EXECUTE FUNCTION decrease_product_amount();

INSERT INTO invoices (invoice_status) VALUES ('closed');
INSERT INTO invoices (invoice_status) VALUES ('opened');

INSERT INTO products (code,description,amount) VALUES ('001', 'Banana', 4);
INSERT INTO products (code,description,amount) VALUES ('002', 'Pera', 5);

INSERT INTO products_to_invoices (invoice_id, product_id) VALUES (1,'001');
INSERT INTO products_to_invoices (invoice_id, product_id) VALUES (1,'002');

INSERT INTO products_to_invoices (invoice_id, product_id) VALUES (2,'001');
INSERT INTO products_to_invoices (invoice_id, product_id) VALUES (2,'002');
