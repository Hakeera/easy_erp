package model

// FichaModelo representa a estrutura dos dados da ficha técnica que será utilizada como base para criação da FichaProduto (Produtos do Banco)
type FichaModelo struct {
	ID             int      `json:"id"`
	Modelo         string   `json:"modelo"`
	Linhas         string   `json:"linhas"`
	Categoria      string   `json:"categoria"`
	Grade          string   `json:"grade"`
	Dimensoes      float64	`json:"dimensoes"`
	Descricao      string   `json:"descricao"`
	Instrucoes     string   `json:"instrucoes"`
	CustoCorte     float64  `json:"custoCorte"`
	CustoCostura   float64  `json:"custoCostura"`
	CustoAcabamento float64 `json:"custoAcabamento"`
	ImagensPaths   []string `json:"imagensPaths"` 
}

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

// Obtém o valor total dos custos de tecidos a partir das informações de cada tecido e da dimensão da ficha modelo utilizada 
func (fp *FichaProduto) ObterCustoTeciedo(modelo FichaModelo) {
	var total float64
	for _, tecido := range fp.Tecidos {
		total += tecido.Custo * tecido.Rendimento / modelo.Dimensoes
	}
	fp.CustoTecido = total
}


