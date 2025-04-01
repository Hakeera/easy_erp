CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    style VARCHAR(50),
    fabric VARCHAR(50),
    color VARCHAR(50),
    size VARCHAR(50),
    client VARCHAR(50),
    line VARCHAR(50),
    category VARCHAR(50),
    situation VARCHAR(50),
    description VARCHAR(500),
    price BIGINT NOT NULL DEFAULT 0,
    technical VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Criar a função para atualizar automaticamente o updated_at
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Criar os triggers para ambas as tabelas
CREATE TRIGGER trigger_update_timestamp_clients
BEFORE UPDATE ON clients
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER trigger_update_timestamp_products
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

