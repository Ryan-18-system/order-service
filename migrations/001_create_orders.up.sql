CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_name TEXT NOT NULL,
    total NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
INSERT INTO orders (customer_name, total) VALUES
  ('João Silva', 199.90),
  ('Maria Oliveira', 349.50),
  ('Carlos Souza', 75.20),
  ('Ana Pereira', 489.00),
  ('Lucas Lima', 129.99);
