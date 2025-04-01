package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/labstack/echo/v4"
)

// ClientRoutes configura as rotas para clientes
func ClientRoutes(e *echo.Echo) {
	e.GET("/clients", controller.GetClients)
	e.POST("/clients", controller.CreateClient) 
	e.PUT("/clients/:id", controller.UpdateClient)
	e.DELETE("/clients/:id", controller.DeleteClient)
}


