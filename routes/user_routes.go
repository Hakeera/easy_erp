package routes

import (
	"github.com/Hakeera/easy_erp/controller"
	"github.com/labstack/echo/v4"
)


func UserRoutes(e *echo.Echo) {

	e.GET("/users", controller.UserPage)
	//e.GET("/users/list", service.GetUsers())
	e.POST("/auth/login", controller.UserLogin)

	e.GET("/users/create", controller.UserForm)
	e.GET("/users/cancel", controller.CancelForm)
	e.POST("/users", controller.CreateUser)

}
