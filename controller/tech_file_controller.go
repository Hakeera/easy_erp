package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Hakeera/easy_erp/model"
	"github.com/labstack/echo/v4"
)

func LoadModelTechFile(c echo.Context) error {
	
	tmpl, err := template.ParseFiles("view/tech_file/fichatech.html",)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar template")
	}
	// Renderiza o template sem passar dados
	return tmpl.Execute(c.Response(), nil)
}

func LoadProductTechFile(c echo.Context) error {
	
	tmpl, err := template.ParseFiles("view/tech_file/fichatec.html",)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar template")
	}
	// Renderiza o template sem passar dados
	return tmpl.Execute(c.Response(), nil)
}


func CriarFichaModelo(c echo.Context) error {
	c.Request().ParseForm() // <-- necessário para popular o PostForm

	grades := c.Request().PostForm["grades_grade"]
	dimensoes := c.Request().PostForm["grades_dimensao"]

	pares := make([]model.ParGradeDimensao, 0, len(grades))

	for i := range grades {
		d, err := strconv.ParseFloat(dimensoes[i], 64)
		if err != nil {
			d = 0
		}
		pares = append(pares, model.ParGradeDimensao{
			Grade:    grades[i],
			Dimensao: d,
		})
	}

	corte, _ := strconv.ParseFloat(c.FormValue("custoCorte"), 64)
	costura, _ := strconv.ParseFloat(c.FormValue("custoCostura"), 64)
	acabamento, _ := strconv.ParseFloat(c.FormValue("custoAcabamento"), 64)

	ficha := model.FichaModelo{
		Modelo:             c.FormValue("modelo"),
		Linhas:             c.FormValue("linhas"),
		Categoria:          c.FormValue("categoria"),
		Descricao:          c.FormValue("descricao"),
		Instrucoes:         c.FormValue("instrucoes"),
		CustoCorte:         corte,
		CustoCostura:       costura,
		CustoAcabamento:    acabamento,
		ParesGradeDimensao: pares,
	}

	// Mostrar no terminal
	fmt.Println("Grades recebidas:", grades)
	fmt.Println("Dimensoes recebidas:", dimensoes)

	// Retornar no frontend
	jsonFicha, err := json.MarshalIndent(ficha, "", "  ")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao gerar JSON")
	}

	return c.HTML(http.StatusOK, "<pre>"+string(jsonFicha)+"</pre>")
}

