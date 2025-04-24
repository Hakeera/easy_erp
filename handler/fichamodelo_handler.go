package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/Hakeera/easy_erp/controller"
)

// CriarFichaModelo repassa a requisição para o controller
func CriarFichaModelo(c echo.Context) error {
	return controller.CriarFichaModelo(c)
}

