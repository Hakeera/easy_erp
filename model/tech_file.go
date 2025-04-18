package model

// FichaProduto representa a estrutura da ficha técnica do Produto final. Ela herda dados da FichaModelo para a criação de cada produto de forma personalizada 
type FichaProduto struct {
	ID              int         `json:"id"`
	FichaModeloID     int       `json:"fichamodeloid"`
	Cliente         string      `json:"cliente"`
	Tecidos         []Tecido    `json:"tecidos"`
	CustoTecido     float64     `json:"custotecido"`
	Aviamentos      []Aviamento `json:"aviamentos"`
	CustoAviamentos float64     `json:"custoaviamentos"`
	Desenhos        []Desenho   `json:"desenhos"`
	CustoDesenho    float64     `json:"custodesenho"`
}

// Obtém o valor total dos custos de aviamentos
func (fp *FichaProduto) ObterCustoAviamento() {
	var total float64
	for _, aviamento := range fp.Aviamentos {
		total += aviamento.CustoUnitario* float64(aviamento.QtdAviamento)
	}
	fp.CustoAviamentos = total
}

