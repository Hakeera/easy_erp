package routes

import (
	"fmt"

	"github.com/Hakeera/easy_erp/controller"
	"github.com/labstack/echo/v4"
)

// SetUpRoutes inicializa todas as rotas da aplicação
// @param e *echo.Echo Instância do framework Echo
func SetUpRoutes(e *echo.Echo)  {
	fmt.Println("Setup routes")
	e.GET("/", controller.HomePage)

	UserRoutes(e)
	ClientRoutes(e)
	TechFileRoutes(e)

}
