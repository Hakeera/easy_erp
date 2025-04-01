package controller

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get Products data and renders products_list.html template
func GetProducts(c echo.Context) error {

	data := struct {
		Title   string
	}{
		Title:   "Página de Produtos",
	}
	// Carregar o template corretamente
	tmpl, err := template.ParseFiles(
		"view/layouts/base.html",
		"view/products/products.html",
	)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar template: " + err.Error())
	}

	// Renderizar o template com os dados
	if err := tmpl.Execute(c.Response().Writer, data); err != nil {
		c.Logger().Error("Erro ao Executar template:", err)
		return c.String(http.StatusInternalServerError, "Erro ao renderizar template")
	}

	return nil
}

func CreateProduct(c echo.Context) error {
	return nil
}

func UpdateProduct(c echo.Context) error {
	return nil
}

func DeleteProduct(c echo.Context) error {
	return nil
}
