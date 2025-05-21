package model

// FichaProduto representa a estrutura da ficha técnica do Produto final. Ela herda dados da FichaModelo para a criação de cada produto de forma personalizada 
type FichaProduto struct {
	ID              int         `json:"id"`
	FichaModeloID   int       `json:"fichamodeloid"`
	Cliente         string      `json:"cliente"`
	Tecidos         []Tecido    `json:"tecidos"`
	CustoTecido     int 	    `json:"custotecido"`
	Aviamentos      []Aviamento `json:"aviamentos"`
	CustoAviamentos int 	    `json:"custoaviamentos"`
	Desenhos        []Desenho   `json:"desenhos"`
	CustoDesenho    int     `json:"custodesenho"`
}
