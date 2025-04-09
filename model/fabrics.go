package model

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
	QtdAviamento  uint    `json:"qtdaviamento"`
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
