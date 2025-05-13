package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/labstack/echo/v4"
)

// ClientRoutes configura as rotas para clientes
func ClientRoutes(e *echo.Echo) {
	e.GET("/clients/home", controller.ClientsPage)
	e.GET("/clients/list", controller.GetClients)
	e.POST("/clients/create", controller.CreateClient) 
	e.PUT("/clients/:id", controller.UpdateClient)
	e.DELETE("/clients/:id", controller.DeleteClient)
}


