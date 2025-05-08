package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)


func HomePage(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/layouts/base.html")

	if err != nil {
		log.Println("Erro ao carregar templates:", err)
		return c.String(http.StatusInternalServerError, "Erro ao carregar templates: "+err.Error())
	}

	// Executa o template base
	return tmpl.ExecuteTemplate(c.Response(), "base.html", nil)
}
