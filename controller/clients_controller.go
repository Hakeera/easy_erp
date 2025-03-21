package controllers

import (
	"net/http"
	"strconv"

	"github.com/Hakeera/easy_erp/configuration"
	models "github.com/Hakeera/easy_erp/model"
	services "github.com/Hakeera/easy_erp/service"
	"github.com/labstack/echo/v4"
)

func CreateClient(c echo.Context) error {
	var client models.Client
	if err := c.Bind(&client); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	if err := services.CreateClient(&client); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, client)
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

// TestDBConnection verifica a conexão com o banco de dados
func TestDBConnection(c echo.Context) error {
	db := configuration.GetDB()

	// Testa a conexão
	if err := db.Exec("SELECT 1").Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Banco de dados inacessível"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Conexão com o banco bem-sucedida!"})
}
