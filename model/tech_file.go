package model

import (
	"time"
)

type Tecido struct {
	ID uint
	Nome string
	FichaTecnica string
	Rendimento string
	Aproveitamento string
	Largura string
	Custo float64
	Observacoes string
}

type Aviamento struct {
	ID            int     `json:"id"`
	Tipo          string  `json:"tipo"`
	Cor           string  `json:"cor"`
	CustoUnitario float64 `json:"custounitario"`
}

type Desenho struct {
	Tipo          string  `json:"tipo"`
	Imagem        string  `json:"imagem"`
	Posicao       string  `json:"posicao"`
	CustoUnitario float64 `json:"custounitario"`
}

type QtdTecido struct {
	Grade		string
	Quantidade	float32
}

type FichaModelo struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Modelo string
	Grade string
	Linhas string
	Categoria string
	Descricao string
	Instrucoes string
	Imagens string

	CustoCostura float64
	CustoCorte float64
	CustoAcabamento float64

	DataCriacao time.Time
	DataAtualiz time.Time

	QtdTecido []QtdTecido

	Status string `gorm:"type:varchar(20);default:'rascunho'" json:"status"` 
}

type FichaProduto struct {
	ID              int         `json:"id"`
	FichaModelo     int         `json:"fichamodelo"`
	Cliente         string      `json:"cliente"`
	Tecidos         []Tecido    `json:"tecidos"`
	CustoTecido     float64     `json:"custotecido"`
	Aviamentos      []Aviamento `json:"aviamentos"`
	CustoAviamentos float64     `json:"custoaviamentos"`
	Desenhos        []Desenho   `json:"desenhos"`
	CustoDesenho    float64     `json:"custodesenho"`
	QtdEstimada     int         `json:"qtdestimada"`
}

type Custos struct {
	ID uint
	Tecido float64
	Aviamento float64
	Corte float64
	Costura float64
	Acabamento float64
	Desenho float64
	Markup float32
}

type Produto struct {
    ID                 uint      `json:"id"`
    FichaModelo     uint      `json:"ficha_modelo"`
    FichaProduto	uint`json:"ficha_produto"`
    ClienteID		uint      `json:"cliente_id"`
    Situacao           string    `json:"situacao"`
    Linha              string    `json:"linha"`
    Categoria 	       string    `json:"categoria"`
    
    ValorVenda         float64   `json:"valor_venda"`

    // Nome 
    Nome string    `json:"nome"`
    Modelo	FichaModelo
    Cliente            Client
    Tecido             Tecido
    Cor string
    Tamanho string  

    // Datas
    DataAtualizacao    time.Time `json:"data_atualizacao"`
}





