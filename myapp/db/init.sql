CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(100)
);

INSERT INTO users (name, email) VALUES
('Taro', 'taro@example.com'),
('Hanako', 'hanako@example.com'),
('Jiro', 'jiro@example.com');