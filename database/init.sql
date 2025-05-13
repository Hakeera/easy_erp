-- Tabela principal de clientes (dados cadastrais básicos)
CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    trading_name VARCHAR(255), -- Nome fantasia (para PJ)
    email VARCHAR(255),
    phone VARCHAR(20),
    address TEXT,
    document_type VARCHAR(10) NOT NULL, -- 'CPF' ou 'CNPJ'
    document_number VARCHAR(20) NOT NULL,
    cep CHAR(8),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT uk_client_document UNIQUE (document_type, document_number)
);

-- Criação da tabela "produtos"
CREATE TABLE IF NOT EXISTS produtos (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(50)
);

