package model

import (
	"fmt"
	"time"
)

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


func (fp *FichaProduto) ObterCustoAviamento() {
	var total float64
	for _, aviamento := range fp.Aviamentos {
		total += aviamento.CustoUnitario* float64(aviamento.QtdAviamento)
	}
	fp.CustoAviamentos = total
}

func main() {
	// Criando um exemplo de ficha de produto
	ficha := FichaProduto{
		ID:          1,
		FichaModelo: 10,
		Cliente:     "Apramed",
		Aviamentos: []Aviamento{
			{ID: 13, Tipo: "Botão", Cor: "Branco", CustoUnitario: 0.1, QtdAviamento: 50},
			{ID: 10, Tipo: "Zíper", Cor: "Azul", CustoUnitario: 0.7, QtdAviamento: 10},
		},
	}

	// Calculando custo dos aviamentos
	ficha.ObterCustoAviamento()

	// Exibindo resultado
	fmt.Printf("Custo total dos aviamentos: %.2f\n", ficha.CustoAviamentos)
}


