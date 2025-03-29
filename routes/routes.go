package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/labstack/echo/v4"
)

// ClientRoutes configura as rotas para clientes
func ClientRoutes(e *echo.Echo) {
	e.GET("/clients", controllers.GetClients)
	e.POST("/clients", controllers.CreateClient) 
	e.PUT("/clients/:id", controllers.UpdateClient)
	e.DELETE("/clients/:id", controllers.DeleteClient)
}

