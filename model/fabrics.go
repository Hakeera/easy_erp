package model

type Tecido struct {
	ID uint
	Nome string
	FichaTecnica string
	Rendimento int 
	Aproveitamento string
	Largura string
	Custo int 
	Observacoes string
}

type Aviamento struct {
	ID            int     `json:"id"`
	Tipo          string  `json:"tipo"`
	Cor           string  `json:"cor"`
	CustoUnitario int`json:"custounitario"`
	QtdAviamento  uint    `json:"qtdaviamento"`
}

type Desenho struct {
	Tipo          string  `json:"tipo"`
	Imagem        string  `json:"imagem"`
	Posicao       string  `json:"posicao"`
	CustoUnitario int`json:"custounitario"`
}

