package controller

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
		"view/layouts/base.html",
		"view/clients/clients.html",
		"view/clients/clients_list.html",
		"view/clients/clients_form.html",
	)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar template: " + err.Error())
	}
	// Dados a serem passados para o template
	data := struct {
		Title   string
		Clients []model.Cliente
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
    log.Println("Iniciando criação de cliente...")

    // Criar um novo cliente com os dados do formulário
    client := models.Cliente{
        Nome:  c.FormValue("name"),
        Email: c.FormValue("email"),
        Telefone: c.FormValue("phone"),
    }

    log.Println("Dados do cliente recebidos:", client)

    // Tentar salvar no banco de dados
    if err := services.CreateClient(&client); err != nil {
        log.Println("Erro ao criar cliente:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao criar cliente"})
    }

    log.Println("Cliente criado com sucesso!")

    // Buscar a lista atualizada de clientes
    log.Println("Buscando lista atualizada de clientes...")
    clients, err := services.GetClients()
    if err != nil {
        log.Println("Erro ao buscar clientes:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao buscar clientes"})
    }

    log.Println("Lista de clientes carregada:", clients)

    // Renderizar o template atualizado
    log.Println("Renderizando template clients_list.html...")
    err = c.Render(http.StatusOK, "clients_list.html", map[string]any{
        "Clients": clients,
    })
    if err != nil {
        log.Println("Erro ao renderizar template:", err)
    }

    log.Println("Template renderizado com sucesso!")
    return nil
}

// Updates Clients by ID
func UpdateClient(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	var updatedClient models.Cliente
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

