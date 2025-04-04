-- Criação da tabela "modelos" (ex: modelos de fichas ou produtos)
CREATE TABLE IF NOT EXISTS modelos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL
);

-- Criação da tabela "clients" (clientes)
CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(50)
);

-- Criação da tabela "tecidos"
CREATE TABLE IF NOT EXISTS tecidos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL
);

-- Criação da tabela "cores"
CREATE TABLE IF NOT EXISTS cores (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL
);

-- Criação da tabela "tamanhos"
CREATE TABLE IF NOT EXISTS tamanhos (
    id SERIAL PRIMARY KEY,
    descricao VARCHAR(50) NOT NULL
);

-- Criação da tabela "fichas_tecnicas"
CREATE TABLE IF NOT EXISTS fichas_tecnicas (
    id SERIAL PRIMARY KEY,
    modelo_id INT REFERENCES modelos(id),
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    instrucoes_producao TEXT,
    imagem_referencia VARCHAR(255),
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Criação da tabela "custos_ficha"
CREATE TABLE IF NOT EXISTS custos_ficha (
    ficha_tecnica_id INT PRIMARY KEY REFERENCES fichas_tecnicas(id),
    tecido DECIMAL(10, 2) DEFAULT 0,
    aviamento DECIMAL(10, 2) DEFAULT 0,
    corte DECIMAL(10, 2) DEFAULT 0,
    costura DECIMAL(10, 2) DEFAULT 0,
    transporte DECIMAL(10, 2) DEFAULT 0,
    embalagem DECIMAL(10, 2) DEFAULT 0,
    desenho DECIMAL(10, 2) DEFAULT 0,
    markup DECIMAL(5, 2) DEFAULT 100.00
);

-- Criação da tabela "tecidos_ficha"
CREATE TABLE IF NOT EXISTS tecidos_ficha (
    id SERIAL PRIMARY KEY,
    ficha_tecnica_id INT REFERENCES fichas_tecnicas(id),
    tecido_id INT REFERENCES tecidos(id),
    quantidade_metros DECIMAL(6, 2),
    observacoes TEXT
);

-- Criação da tabela "aviamentos_ficha"
CREATE TABLE IF NOT EXISTS aviamentos_ficha (
    id SERIAL PRIMARY KEY,
    ficha_tecnica_id INT REFERENCES fichas_tecnicas(id),
    nome VARCHAR(100),
    quantidade INT,
    custo_unitario DECIMAL(10, 2),
    observacoes TEXT
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

