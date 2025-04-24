package model

import (
	"gorm.io/datatypes"
)


type FichaModeloDB struct {
	ID    uint           `gorm:"primaryKey"`
	Dados datatypes.JSON `gorm:"type:jsonb"`
}

// Força o GORM a usar o nome `fichas_modelo`
func (FichaModeloDB) TableName() string {
	return "fichas_modelo"
}

type ParGradeDimensao struct {
	Grade    string  `json:"grade"`
	Dimensao float64 `json:"dimensao"`
}

// FichaModelo representa a estrutura dos dados da ficha técnica que será utilizada como base para criação da FichaProduto (Produtos do Banco)
type FichaModelo struct {
	Modelo           string              `json:"modelo"`
	Linhas           string              `json:"linhas"`
	Categoria        string              `json:"categoria"`
	Descricao        string              `json:"descricao"`
	Instrucoes       string              `json:"instrucoes"`
	CustoCorte       float64             `json:"custoCorte"`
	CustoCostura     float64             `json:"custoCostura"`
	CustoAcabamento  float64             `json:"custoAcabamento"`
	ParesGradeDimensao []ParGradeDimensao `json:"paresGradeDimensao"`
}

