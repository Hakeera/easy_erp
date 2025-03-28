package controllers

import (
	"html/template"
	"net/http"

	"github.com/Hakeera/easy_erp/model"
	"github.com/Hakeera/easy_erp/service"
	"github.com/labstack/echo/v4"
)

// GetHome renderiza base.html
func GetHome(c echo.Context) error {
	// Buscar a lista de clientes
	clients, err := service.GetClients()
	if err != nil {
		c.Logger().Error("Erro ao buscar clientes:", err)
		return c.String(http.StatusInternalServerError, "Erro ao buscar clientes")
	}

	// Carregar o template corretamente
	tmpl, err := template.ParseFiles(
		"view/base.html",
		"view/clients.html",
		"view/clients_list.html",
		"view/clients_form.html",
	)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar template: " + err.Error())
	}
		// Dados a serem passados para o template
	data := struct {
		Title   string
		Clients []model.Client
	}{
		Title:   "Página de Clientes",
		Clients: clients,
	}

	// Renderizar o template com os dados
	if err := tmpl.Execute(c.Response().Writer, data); err != nil {
		c.Logger().Error("Erro ao renderizar template:", err)
		return c.String(http.StatusInternalServerError, "Erro ao renderizar template")
	}

	return nil

}


