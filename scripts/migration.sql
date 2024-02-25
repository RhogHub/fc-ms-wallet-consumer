CREATE DATABASE IF NOT EXISTS wallet;
USE wallet;

CREATE TABLE IF NOT EXISTS clients (
    id VARCHAR(255),
    name VARCHAR(255),
    email VARCHAR(255),
    created_at date
);

INSERT INTO clients (id, name, email, created_at) VALUES
('da5151aa-8c22-455c-aae0-fd7d5b2b31c4','John Doe', 'john@example.com', '2024-02-18'),
('bc5823ad-6d17-404d-b13b-9150efd12b3b','Jane Smith', 'jane@example.com', '2024-02-18');

CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(255),
    client_id VARCHAR(255),
    balance int NOT NULL DEFAULT 0,
    created_at date
);

INSERT INTO accounts (id, client_id, balance, created_at) VALUES
('3351514a-8c22-455c-aae0-fd7d5b2b31j7','da5151aa-8c22-455c-aae0-fd7d5b2b31c4', 1000, '2024-02-18'),
('023b19f6-c81f-497a-bdc2-4602e7856632','bc5823ad-6d17-404d-b13b-9150efd12b3b', 500, '2024-02-18');

CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255),
    account_id_from VARCHAR(255),
    account_id_to varchar(255),
    amount int NOT NULL DEFAULT 0,
    created_at date
);

CREATE DATABASE IF NOT EXISTS db_consumer;
USE db_consumer;

CREATE TABLE IF NOT EXISTS balances (
    id VARCHAR(255),
    account_id VARCHAR(255),
    balance decimal(10,2) NOT NULL DEFAULT 0.00,
    created_at date,
    updated_at date
);

-- INSERT INTO balances (id, account_id, balance, created_at, updated_at) VALUES
-- ('9951514a-8c29-455c-aae0-fd7d5b2b3h99','j65151aa-8c23-655c-bae4-fd7d5b2b31d3', 1000.00, '2024-02-25', '2024-02-25');

SET GLOBAL sql_mode = '';