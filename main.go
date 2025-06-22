package main

import (
	"html/template"
	"io"
	"log"
	"os"

	"github.com/Hakeera/easy_erp/config"
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

	// TESTE: Verificar se o banco está funcionando
	db := config.GetDB()
	if db != nil {
		log.Println("✅ Banco de dados conectado com sucesso!")

		// Teste simples de conexão
		sqlDB, err := db.DB()
		if err == nil {
			if err := sqlDB.Ping(); err == nil {
				log.Println("✅ Ping no banco OK!")
			} else {
				log.Printf("❌ Erro no ping: %v", err)
			}
		}
	} else {
		log.Println("❌ Banco de dados é nil!")
	}

	e := echo.New()

	// Configura o renderer de templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("view/**/*.html")),
	}

	e.Renderer = renderer

	// Configurar rotas
	routes.SetUpRoutes(e)

	// Iniciar o servidor
	e.Logger.Fatal(e.Start(":1323"))
}

