package main

import (
	"html/template"
	"io"
	"log"
	"os"

	config "github.com/Hakeera/easy_erp/configuration"
	"github.com/Hakeera/easy_erp/routes"
	"github.com/labstack/echo/v4"
)

// TemplateRenderer para Echo
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {

	config.LoadEnv() 

    	log.Println("DB_HOST:", os.Getenv("DB_HOST"))

	// Inicializa o banco de dados
	config.InitDB()

	e := echo.New()
	
	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("view/**/*.html")), 
}
e.Renderer = t
	// Configurar rotas
	routes.SetUpRoutes(e)

	// Iniciar o servidor
	e.Logger.Fatal(e.Start(":1323"))
}

