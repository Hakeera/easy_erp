package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/labstack/echo/v4"
)


func TechFileRoutes(e *echo.Echo)  {

	e.GET("/tech_file/product", controller.LoadProductTechFile)
	e.POST("/fichamodelo", controller.CriarFichaModelo)

}
