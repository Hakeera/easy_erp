package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	models "github.com/Hakeera/easy_erp/model"
	services "github.com/Hakeera/easy_erp/service"
	"github.com/labstack/echo/v4"
)

func ListClients(c echo.Context) error {
	// Carregar o template corretamente
	tmpl, err := template.ParseFiles("view/clients_list.html") 
	if err != nil {
		return fmt.Errorf("erro ao carregar o template: %w", err)
	}

	// Obter a lista de clientes do banco
	clients, err := services.GetClients()
	if err != nil {
		return fmt.Errorf("erro ao buscar clientes: %w", err)
	}

	// Criar os dados a serem passados para o template
	data := struct {
		Title   string
		Clients []models.Client
	}{
		Title:   "Lista de Clientes",
		Clients: clients,
	}

	// Definir o tipo de resposta como HTML antes de executar o template
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")

	// Executar o template e renderizar para o usuário
	if err := tmpl.Execute(c.Response().Writer, data); err != nil {
		return fmt.Errorf("erro ao renderizar template: %w", err)
	}

	return nil
}

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
        return c.String(http.StatusInternalServerError, "Erro ao buscar clientes")
    }

    // Retornar apenas a lista atualizada para ser injetada no HTML via HTMX
    return c.Render(http.StatusOK, "base.html", map[string]any{
        "Clients": clients,
    })
}

func GetClients(c echo.Context) error {
	clients, err := services.GetClients()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, clients)
}

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

// Renderiza a página de clientes
func GetClientsPage(c echo.Context) error {
	clients, err := services.GetClients()
	if err != nil {
		log.Println("Erro ao buscar clientes:", err)
		return c.String(http.StatusInternalServerError, "Erro ao carregar clientes")
	}

	log.Println("Clientes encontrados:", clients) 
return c.Render(http.StatusOK, "clients.html", map[string]any{
	"Clients": clients,
})

}
