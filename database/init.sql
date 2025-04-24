-- Criação da tabela "clients" (clientes)
CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(50)
);

-- Criação da tabela "produtos"
CREATE TABLE IF NOT EXISTS produtos (
    id SERIAL PRIMARY KEY,
    ficha_tecnica_id INT REFERENCES fichas_tecnicas(id),
    cliente_id INT REFERENCES clients(id),
    tecido_id INT REFERENCES tecidos(id),
    cor_id INT REFERENCES cores(id),
    tamanho_id INT REFERENCES tamanhos(id),
    quantidade INT DEFAULT 1,
    linha VARCHAR(50),
    situacao VARCHAR(50) NOT NULL DEFAULT 'Em Produção',
    custo_calculado DECIMAL(10, 2) NOT NULL,
    valor_venda DECIMAL(10, 2) NOT NULL,
    desconto_percentual DECIMAL(5, 2) DEFAULT 0,
    personalizacao TEXT,
    nome_gerado VARCHAR(255),
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

