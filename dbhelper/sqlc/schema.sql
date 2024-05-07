CREATE TABLE IF NOT EXISTS employee (
    id SERIAL PRIMARY KEY,
    employee_name VARCHAR(100) UNIQUE NOT NULL,
    position VARCHAR(100) UNIQUE NOT NULL,
    salary float NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);