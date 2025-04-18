package model

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

