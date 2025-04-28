package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/Hakeera/easy_erp/handler"
	"github.com/labstack/echo/v4"
)

// TechFileRoutes configura as rotas relacionadas às fichas técnicas de modelo
// @param e *echo.Echo Instância do framework Echo
func TechFileRoutes(e *echo.Echo)  {
	e.GET("/fichamodelo", handler.PaginaFichaModelo)
	e.GET("/fichamodelo/list", controller.ListarFichasModelo)
	e.POST("/fichamodelo/create", handler.CriarFichaModelo)
}
