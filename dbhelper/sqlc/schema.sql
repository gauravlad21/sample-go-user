CREATE TABLE IF NOT EXISTS employee (
    id SERIAL PRIMARY KEY,
    employee_name VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL,
    salary float NOT NULL,
    version int DEFAULT 1,
    created TIMESTAMP DEFAULT current_timestamp,
    updated TIMESTAMP DEFAULT current_timestamp
);