package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/labstack/echo/v4"
)

// ClientRoutes configura as rotas para clientes
func ProductsRoutes(e *echo.Echo) {
	e.GET("/products", controller.GetProducts)
	e.POST("/products", controller.CreateProduct) 
	e.PUT("/products/:id", controller.UpdateProduct)
	e.DELETE("/products/:id", controller.DeleteProduct)
}

