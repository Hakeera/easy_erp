package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/Hakeera/easy_erp/handler"
	"github.com/labstack/echo/v4"
)


func TechFileRoutes(e *echo.Echo)  {

	e.GET("/tech_file/model", controller.LoadProductTechFile)
	e.POST("/fichamodelo", handler.CriarFichaModelo)

}
