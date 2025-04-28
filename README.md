# Easy_ERP

Sistema de ERP especializado para empresas de confecção de vestuário.

## Sobre o Projeto

Easy_ERP é uma solução completa para gerenciamento de recursos empresariais focada no setor de confecções. O sistema permite gerenciar todo o ciclo produtivo, desde a criação de fichas técnicas até o controle de produção e vendas.

## Estrutura do Projeto

```
easy_erp/
├── config/         # Configurações do sistema e banco de dados
├── controller/     # Controladores das funcionalidades
├── handler/        # Manipuladores HTTP para as rotas
├── model/          # Definição dos modelos de dados
├── routes/         # Configuração das rotas da API
└── view/           # Arquivos de interface do usuário
```

## Funcionalidades Principais

- **Fichas Técnicas de Modelo**: Criação e gestão de fichas técnicas para produtos, incluindo:
  - Informações do modelo
  - Categorização de produtos
  - Estrutura de grades e dimensões
  - Cálculo de custos (corte, costura, acabamento)

## Tecnologias Utilizadas

- **Backend**: Go (Golang)
- **Framework Web**: Echo
- **ORM**: GORM
- **Banco de Dados**: PostgreSQL
- **Frontend**: HTML, HTMX, CSS, JavaScript

## Requisitos

- Go 1.16+
- PostgreSQL 12+
- Variáveis de ambiente configuradas (ver seção de configuração)

## Configuração

### Variáveis de Ambiente

O sistema utiliza variáveis de ambiente para configuração. Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=easy_erp_db
DB_HOST=localhost
DB_PORT=5432
```

## Instalação e Execução

1. Clone o repositório:
   ```bash
   git clone https://github.com/Hakeera/easy_erp.git
   cd easy_erp
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Configure o banco de dados:
   ```bash
   # Certifique-se que o PostgreSQL está rodando
   # Crie o banco de dados "easy_erp_db"
   ```

4. Inicie a aplicação:
   ```bash
   go run main.go
   ```

5. Acesse a aplicação em `http://localhost:1323`

## Arquitetura

O projeto segue uma arquitetura em camadas:

1. **Rotas**: Definição dos endpoints da API (routes/)
2. **Handlers**: Processamento inicial das requisições (handler/)
3. **Controllers**: Lógica de negócio (controller/)
4. **Models**: Estruturas de dados e acesso ao banco (model/)
5. **Config**: Configurações do sistema (config/)

## Modelo de Dados

### Ficha Técnica de Modelo

A ficha técnica é o documento base para a produção das peças de vestuário, contendo:

- Identificação do modelo
- Descrição e instruções
- Categoria e linha de produto
- Grades e dimensões (tamanhos)
- Custos de produção segregados (corte, costura, acabamento)

Os dados são armazenados em formato JSON no banco de dados para maior flexibilidade.

## Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova funcionalidade'`)
4. Push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).

---
