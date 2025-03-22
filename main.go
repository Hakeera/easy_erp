package main

import (
	"log"
	"os"

	"github.com/Hakeera/easy_erp/configuration"
	"github.com/Hakeera/easy_erp/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	configuration.LoadEnv() 

    	log.Println("DB_HOST:", os.Getenv("DB_HOST"))



	// Inicializa o banco de dados
	configuration.InitDB()

	e := echo.New()

	// Configurar rotas
	routes.ClientRoutes(e)

	// Iniciar o servidor
	e.Logger.Fatal(e.Start(":1323"))
}

