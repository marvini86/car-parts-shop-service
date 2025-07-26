CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    code_integration VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    image VARCHAR(50) NOT NULL,
    available_quantity INT NOT NULL,
    category_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS item_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status VARCHAR(50) NOT NULL,
    total_value DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    item_id INT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS order_payment_details (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    card_number VARCHAR(50) NOT NULL,
    expiry_date VARCHAR(50) NOT NULL,
    cvv VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS order_delivery_addresses (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    address VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    state VARCHAR(50) NOT NULL,
    zip_code VARCHAR(50) NOT NULL,
    country VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE orders ADD CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE order_items ADD CONSTRAINT order_items_order_id_fkey FOREIGN KEY (order_id) REFERENCES orders(id);
ALTER TABLE order_items ADD CONSTRAINT order_items_item_id_fkey FOREIGN KEY (item_id) REFERENCES items(id);
ALTER TABLE order_payment_details ADD CONSTRAINT payment_details_order_id_fkey FOREIGN KEY (order_id) REFERENCES orders(id);
ALTER TABLE order_delivery_addresses ADD CONSTRAINT order_delivery_addresses_order_id_fkey FOREIGN KEY (order_id) REFERENCES orders(id);

ALTER TABLE items ADD CONSTRAINT items_category_id_fkey FOREIGN KEY (category_id) REFERENCES item_categories(id);

INSERT INTO item_categories (name, created_at, updated_at) VALUES ('Car parts', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO item_categories (name, created_at, updated_at) VALUES ('Motorcycle parts', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO item_categories (name, created_at, updated_at) VALUES ('Truck parts', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO item_categories (name, created_at, updated_at) VALUES ('Bus parts', '2023-01-01 00:00:00', '2023-01-01 00:00:00');

INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000001', 'Tire', 'Replacement', '100', 'https://via.placeholder.com/150', '1', '4', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000002', 'Brake', 'Replacement', '100', 'https://via.placeholder.com/150', '1', '3', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000003', 'Engine', 'Replacement', '100', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000004', 'Transmission', 'Replacement', '100', 'https://via.placeholder.com/150', '1', '2', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000005', 'Battery', 'Replacement', '100', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000006', 'Oil', 'Change', '100', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000007', 'Tire', 'Change', '100', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000008', 'Brake', 'Change', '100', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000009', 'Engine', 'Change', '100', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000010', 'Transmission', 'Change', '100', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000011', 'Battery', 'Change', '100', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');
INSERT INTO items (code_integration, name, description, price, image, category_id, available_quantity, created_at, updated_at) VALUES ('000000012', 'Oil', 'Change', '100', 'https://via.placeholder.com/150', '1', '10', '2023-01-01 00:00:00', '2023-01-01 00:00:00');

INSERT INTO users (name, email, password, created_at, updated_at) VALUES ('Marcus Vinicius', 'marcusvinicius@gmail.com', '123456', '2023-01-01 00:00:00', '2023-01-01 00:00:00');


