package model

import "time"

type Produto struct {
    ID                 uint      `json:"id"`
    FichaModelo     	uint      `json:"ficha_modelo"`
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
