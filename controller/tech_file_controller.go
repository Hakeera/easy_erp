package controller

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoadModelTechFile(c echo.Context) error {
	
	tmpl, err := template.ParseFiles("view/tech_file/ficha_tecnica_modelo_form.html",)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar template")
	}
	// Renderiza o template sem passar dados
	return tmpl.Execute(c.Response(), nil)
}

func LoadProductTechFile(c echo.Context) error {
	
	tmpl, err := template.ParseFiles("view/tech_file/ficha_tecnica_produto_form.html",)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar template")
	}
	// Renderiza o template sem passar dados
	return tmpl.Execute(c.Response(), nil)
}
