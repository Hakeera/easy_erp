package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/Hakeera/easy_erp/controller"
)

// PaginaFichaModelo redireciona a requisição para o controlador correspondente
// @param c echo.Context Contexto da requisição
// @return error Erro de processamento, se houver
func PaginaFichaModelo(c echo.Context) error {
	return controller.PaginaFichaModelo(c)
}

// ListarFichasModelo redireciona a requisição para o controlador correspondente
// @param c echo.Context Contexto da requisição
// @return error Erro de processamento, se houver
func ListarFichasModelo(c echo.Context) error {
	return controller.ListarFichasModelo(c) 
}

// CriarFichaModelo redireciona a requisição para o controlador correspondente
// @param c echo.Context Contexto da requisição
// @return error Erro de processamento, se houver
func CriarFichaModelo(c echo.Context) error {
	return controller.CriarFichaModelo(c)
}
