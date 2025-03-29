package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Hakeera/easy_erp/model"
	models "github.com/Hakeera/easy_erp/model"
	"github.com/Hakeera/easy_erp/service"
	services "github.com/Hakeera/easy_erp/service"
	"github.com/labstack/echo/v4"
)

// Get Clients data and renders clients_list.html template
func GetClients(c echo.Context) error {
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
// Creates a Client and renders template client_list.html 
func CreateClient(c echo.Context) error {
    client := models.Client{
        Name:  c.FormValue("name"),
        Email: c.FormValue("email"),
        Phone: c.FormValue("phone"),
    }

    if err := services.CreateClient(&client); err != nil {
        return c.String(http.StatusInternalServerError, "Erro ao criar cliente")
    }

    // Buscar a lista atualizada de clientes
    clients, err := services.GetClients()
    if err != nil {
	    log.Println("Erro ao Criar Cliente:", err)
        return c.String(http.StatusInternalServerError, "Erro ao buscar clientes")
    }

    // Retornar apenas a lista atualizada para ser injetada no HTML via HTMX
    return c.Render(http.StatusOK, "clients_list.html", map[string]any{
        "Clients": clients,
    })
}

// Updates Clients by ID
func UpdateClient(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	var updatedClient models.Client
	if err := c.Bind(&updatedClient); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	err = services.UpdateClient(uint(id), &updatedClient)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cliente atualizado com sucesso"})
}

func DeleteClient(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	err = services.DeleteClient(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cliente excluído com sucesso"})
}

