package model

import "time"

type Produto struct {
    ID                 uint      `json:"id"`
    FichaModelo        uint      `json:"ficha_modelo"`
    FichaProduto       uint	 `json:"ficha_produto"`
    ClienteID	       uint      `json:"cliente_id"`
    Situacao           string    `json:"situacao"`
    Linha              string    `json:"linha"`
    Categoria 	       string    `json:"categoria"`
    
    ValorVenda          int	 `json:"valor_venda"`

    // Nome 
    Nome string         `json:"nome"`
    Modelo		FichaModelo
    Cliente             Client
    Tecido              Tecido
    Cor 		string
    Tamanho 		string  

    // Datas
    DataAtualizacao    time.Time `json:"data_atualizacao"`
}

type Custos struct {
	ID	   uint
	Tecido	   int 
	Aviamento  int 
	Corte	   int 
	Costura	   int 
	Acabamento int 
	Desenho    int 
	Markup     float32 
}
