package model

import (
	"gorm.io/datatypes"
)

// FichaModeloDB representa o modelo de dados para armazenamento no banco
// Os dados da ficha são armazenados como JSON.
type FichaModeloDB struct {
	ID    uint           `gorm:"primaryKey"`
	Dados datatypes.JSON `gorm:"type:jsonb"`
}

// TableName sobrescreve o nome da tabela no banco de dados
// @return string Nome da tabela no banco de dados
func (FichaModeloDB) TableName() string {
	return "fichas_modelo"
}

// ParGradeDimensao representa um par de grade e dimensão associada
// usado na configuração de tamanhos das peças
type ParGradeDimensao struct {
	Grade    string  `json:"grade"`    // Identificador da grade (ex: "P", "M", "G")
	Dimensao float64 `json:"dimensao"` // Valor numérico associado à grade
}

// FichaModelo representa a estrutura dos dados da ficha técnica
// que será utilizada como base para criação da FichaProduto (Produtos do Banco)
type FichaModelo struct {
	ID                 uint              `json:"ID"`
	Modelo             string            `json:"modelo"`      // Identificador do modelo
	Linhas             string            `json:"linhas"`      // Linhas de produto
	Categoria          string            `json:"categoria"`   // Categoria do produto
	Descricao          string            `json:"descricao"`   // Descrição detalhada
	Instrucoes         string            `json:"instrucoes"`  // Instruções especiais
	CustoCorte         float64           `json:"custoCorte"`  // Custo do processo de corte
	CustoCostura       float64           `json:"custoCostura"`// Custo do processo de costura
	CustoAcabamento    float64           `json:"custoAcabamento"` // Custo de acabamento
	ParesGradeDimensao []ParGradeDimensao `json:"paresGradeDimensao"` // Configuração de grades e dimensões
}
