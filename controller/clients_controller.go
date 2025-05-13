package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Hakeera/easy_erp/model"
	models "github.com/Hakeera/easy_erp/model"
	"github.com/Hakeera/easy_erp/service"
	"github.com/labstack/echo/v4"
)

// ClientsPage renderiza a página de clientes com os dados carregados do banco
func ClientsPage(c echo.Context) error {
	// Busca os clientes
	clients, err := service.GetClients()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao buscar clientes: " + err.Error())
	}

	// Carrega os templates
	tmpl, err := template.ParseFiles(
		"view/layouts/base.html",
		"view/clients/clients.html",
		"view/clients/clients_form.html",
		"view/clients/clients_list.html",
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar templates: " + err.Error())
	}

	// Prepara os dados
	data := map[string]any{
		"Clients": clients,
	}

	// Renderiza o template base com os dados
	if err := tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data); err != nil {
		c.Logger().Error("Erro ao renderizar template:", err)
		return c.String(http.StatusInternalServerError, "Erro ao renderizar template")
	}

	return nil
}

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
		"view/clients/clients_list.html",
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
    log.Println("Iniciando criação de cliente...")
    
    client := models.Client{
        Name:           c.FormValue("name"),
        TradingName:    c.FormValue("trading_name"),
        Email:          c.FormValue("email"),
        Phone:          c.FormValue("phone"),
        Address:        c.FormValue("address"),
        DocumentType:   c.FormValue("document_type"),
        DocumentNumber: c.FormValue("document_number"),
        CEP:            c.FormValue("cep"),
    }

    log.Println("Dados do cliente recebidos:", client)

    // Validação
    if client.Name == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nome é obrigatório"})
    }
    if client.DocumentType == "" || client.DocumentNumber == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Tipo e número do documento são obrigatórios"})
    }

    documentNumberClean := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(client.DocumentNumber, ".", ""), "-", ""), "/", "")
    if client.DocumentType == "CPF" && len(documentNumberClean) != 11 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "CPF inválido"})
    }
    if client.DocumentType == "CNPJ" && len(documentNumberClean) != 14 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "CNPJ inválido"})
    }

    client.DocumentNumber = documentNumberClean
    client.CEP = strings.ReplaceAll(client.CEP, "-", "")

    if err := service.CreateClient(&client); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    clients, err := service.GetClients()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao buscar clientes"})
    }

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

	err = service.UpdateClient(uint(id), &updatedClient)
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

	err = service.DeleteClient(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cliente excluído com sucesso"})
}

