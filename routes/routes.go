package routes

import "github.com/labstack/echo/v4"

// SetUpRoutes inicializa todas as rotas da aplicação
// @param e *echo.Echo Instância do framework Echo
func SetUpRoutes(e *echo.Echo)  {
	ClientRoutes(e)
	TechFileRoutes(e)
}
