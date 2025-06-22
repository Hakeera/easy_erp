package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Hakeera/easy_erp/config"
	"github.com/Hakeera/easy_erp/model"
	"github.com/Hakeera/easy_erp/service"
	"github.com/labstack/echo/v4"
)

// UserPage - Renderiza a página principal de usuários
func UserPage(c echo.Context) error {
	
	tmpl, err := template.ParseFiles(
		"view/user/user_page.html",
		"view/layouts/base.html",
	)
	if err != nil {
		log.Println("Erro ao carregar templates:", err)
		return c.String(http.StatusInternalServerError, "Erro ao carregar User Template: "+err.Error())
	}
	
	data := map[string]any{
		"Title": "Gerenciamento de Usuários - Easy ERP",
	}
	
	return tmpl.ExecuteTemplate(c.Response(), "base.html", data)
}

// Carrega o Formulário de Usuário
func UserForm(c echo.Context) error {
    tmpl, err := template.ParseFiles("view/user/user_form.html")
    if err != nil {
        log.Println("Erro ao carregar template:", err)
        return c.String(http.StatusInternalServerError, "Erro ao carregar template: "+err.Error())
    }
    
    // Executa apenas o template content do formulário
    return tmpl.ExecuteTemplate(c.Response(), "user_form", nil)
}

func CreateUser(c echo.Context) error {
    // Recovery para evitar crash do servidor
    defer func() {
        if r := recover(); r != nil {
            log.Printf("PANIC CAPTURADO em CreateUser: %v", r)
            renderError(c, "Erro interno do servidor")
        }
    }()

    // Log para debug
    log.Printf("CreateUser chamado - Method: %s, Content-Type: %s", 
        c.Request().Method, c.Request().Header.Get("Content-Type"))
    
    // Verificar se é POST
    if c.Request().Method != http.MethodPost {
        return renderError(c, "Método não permitido")
    }

    user := &model.User{
        Username: c.FormValue("username"),
        Email:    c.FormValue("email"),
        Password: c.FormValue("password"),
    }

    // Log dos valores recebidos (sem a senha)
    log.Printf("Dados recebidos - Username: '%s', Email: '%s'", 
        user.Username, user.Email)

    // Validações básicas no controller
    if user.Username == "" || user.Email == "" || user.Password == "" {
        log.Printf("Validação falhou - campos vazios")
        return renderError(c, "Todos os campos são obrigatórios")
    }

    // Chamar o service
    log.Printf("Chamando service.CreateUser...")
    createdUser, err := service.CreateUser(user)
    if err != nil {
        log.Printf("Erro no service: %v", err)
        return renderError(c, "Erro ao criar usuário: "+err.Error())
    }

    log.Printf("Usuário criado com sucesso: %s", createdUser.Username)
    return renderSuccess(c, "Usuário '"+createdUser.Username+"' criado com sucesso!")
}

func UserLogin(c echo.Context)  error {
    // Recovery para evitar crash do servidor
    defer func() {
        if r := recover(); r != nil {
            log.Printf("PANIC CAPTURADO em CreateUser: %v", r)
            renderError(c, "Erro interno do servidor")
        }
    }()

    db := config.GetDB()

    log.Printf("UserLogin chamado - Method: %s, Content-Type: %s", 
        c.Request().Method, c.Request().Header.Get("Content-Type"))
    
    // Verificar se é POST
    if c.Request().Method != http.MethodGet{
        return renderError(c, "Método não permitido")
    }

    user := &model.User{
        Username: c.FormValue("username"),
        Email:    c.FormValue("email"),
        Password: c.FormValue("password"),
    }
}

// Função renderError corrigida
func renderError(c echo.Context, message string) error {
    log.Printf("Renderizando erro: %s", message)
    
    tmpl, err := template.ParseFiles("view/user/user_error.html")
    if err != nil {
        log.Printf("Erro ao carregar template de erro: %v", err)
        return c.String(http.StatusInternalServerError, "Erro ao carregar template")
    }
    
    // Definir headers explicitamente
    c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
    c.Response().Header().Set("Cache-Control", "no-cache")
    
    return tmpl.ExecuteTemplate(c.Response(), "user_error", message)
}

// Função renderSuccess corrigida
func renderSuccess(c echo.Context, message string) error {
    log.Printf("Renderizando sucesso: %s", message)
    
    tmpl, err := template.ParseFiles("view/user/user_success.html")
    if err != nil {
        log.Printf("Erro ao carregar template de sucesso: %v", err)
        return c.String(http.StatusInternalServerError, "Erro ao carregar template")
    }
    
    // Definir headers explicitamente
    c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
    c.Response().Header().Set("Cache-Control", "no-cache")
    
    return tmpl.ExecuteTemplate(c.Response(), "user_success", message)
}


func CancelForm(c echo.Context) error {
    loginHTML := `
    <div class="login-form">
        <h3>Login</h3>
        <form hx-post="/auth/login" hx-target="#auth-status" hx-swap="innerHTML">
            <div class="form-group">
                <label for="login-email">Email:</label>
                <input type="email" id="login-email" name="email" required>
            </div>
            <div class="form-group">
                <label for="login-password">Senha:</label>
                <input type="password" id="login-password" name="password" required>
            </div>
            <button type="submit">Entrar</button>
            <button hx-get="/users/create" hx-target="#auth-status" hx-swap="innerHTML">Criar Novo</button>
        </form>
    </div>`
    
    return c.HTML(http.StatusOK, loginHTML)
}
