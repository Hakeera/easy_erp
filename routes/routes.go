package routes

import "github.com/labstack/echo/v4"

// Inicializes all routes
func SetUpRoutes(e *echo.Echo)  {
	ClientRoutes(e)
	ProductsRoutes(e)
}
