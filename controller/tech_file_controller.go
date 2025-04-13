package controller

import (
	"html/template"
	"net/http"

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

// SaveModelTechFile processa os dados do formulário de ficha técnica de modelo
func SaveModelTechFile(c echo.Context) error {
	
	return nil	
} 
