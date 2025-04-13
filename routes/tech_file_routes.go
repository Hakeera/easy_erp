package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/labstack/echo/v4"
)


func TechFileRoutes(e *echo.Echo)  {

	e.GET("/tech_file/model", controller.LoadModelTechFile)
	e.POST("/tech_file/model/save", controller.SaveModelTechFile)
	e.GET("/tech_file/product", controller.LoadProductTechFile)

	e.POST("/tech_file/model/save", controller.SaveModelTechFile)

}
