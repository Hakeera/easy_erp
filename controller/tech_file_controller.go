package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	config "github.com/Hakeera/easy_erp/configuration"
	"github.com/Hakeera/easy_erp/model"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
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

// Permanencia no banco de dados
func CriarFichaModelo(c echo.Context) error {
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

	custoCorte, _ := strconv.ParseFloat(c.FormValue("custoCorte"), 64)
	custoCostura, _ := strconv.ParseFloat(c.FormValue("custoCostura"), 64)
	custoAcabamento, _ := strconv.ParseFloat(c.FormValue("custoAcabamento"), 64)

	// Struct lógica (sem GORM)
	ficha := model.FichaModelo{
		Modelo:             c.FormValue("modelo"),
		Linhas:             c.FormValue("linhas"),
		Categoria:          c.FormValue("categoria"),
		Descricao:          c.FormValue("descricao"),
		Instrucoes:         c.FormValue("instrucoes"),
		CustoCorte:         custoCorte,
		CustoCostura:       custoCostura,
		CustoAcabamento:    custoAcabamento,
		ParesGradeDimensao: pares,
	}

	// Serializa para JSON
	jsonData, err := json.Marshal(ficha)
	if err != nil {
		log.Println("Erro ao serializar ficha:", err)
		return c.String(http.StatusInternalServerError, "Erro ao processar ficha")
	}

	// Struct de persistência com GORM
	fichaDB := model.FichaModeloDB{
		Dados: datatypes.JSON(jsonData),
	}

	// Salva no banco com GORM
	if err := config.GetDB().Create(&fichaDB).Error; err != nil {
		log.Println("Erro ao salvar ficha:", err)
		return c.String(http.StatusInternalServerError, "Erro ao salvar ficha")
	}

	return c.String(http.StatusOK, "Ficha Técnica salva com sucesso!")
}


