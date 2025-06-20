package controller

import (
	"html/template"
	"log"
	"net/http"

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

// Tratamentos de dados básicos e chama service para o Banco de Dados  
func CreateUser(c echo.Context) error {
    // Log para debug
    log.Printf("CreateUser chamado - Method: %s, Content-Type: %s", c.Request().Method, c.Request().Header.Get("Content-Type"))
    
    // Verificar se é POST
    if c.Request().Method != http.MethodPost {
        return c.String(http.StatusMethodNotAllowed, "Método não permitido")
    }

    user := &model.User{
        Username: c.FormValue("username"),
        Email:    c.FormValue("email"),
        Password: c.FormValue("password"),
    }

    // Log dos valores recebidos
    log.Printf("Dados recebidos - Username: %s, Email: %s, Password: %s", 
        user.Username, user.Email, "***")

    // Validações básicas
    if user.Username == "" || user.Email == "" || user.Password == "" {
        log.Printf("Validação falhou - campos vazios")
        return renderError(c, "Todos os campos são obrigatórios")
    }

    // Chamar o service
    createdUser, err := service.CreateUser(user)
    if err != nil {
        log.Printf("Erro no service: %v", err)
        return renderError(c, "Erro ao criar usuário: "+err.Error())
    }

    log.Printf("Usuário criado com sucesso: %s", createdUser.Username)
    return renderSuccess(c, "Usuário '"+createdUser.Username+"' criado com sucesso!")
}


// Helpers para renderizar mensagens
func renderSuccess(c echo.Context, message string) error {
	tmpl, err := template.ParseFiles("view/user/user_success.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar template")
	}
	
	return tmpl.ExecuteTemplate(c.Response(), "user_success", message)
}

func renderError(c echo.Context, message string) error {
	tmpl, err := template.ParseFiles("view/user/user_error.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erro ao carregar template")
	}
	
	c.Response().WriteHeader(http.StatusBadRequest)
	return tmpl.ExecuteTemplate(c.Response(), "user_error", message)
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
