package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/labstack/echo/v4"
)

// SetUpRoutes inicializa todas as rotas da aplicação
// @param e *echo.Echo Instância do framework Echo
func SetUpRoutes(e *echo.Echo)  {
	e.GET("/", controller.HomePage)
	ClientRoutes(e)
	TechFileRoutes(e)
}
