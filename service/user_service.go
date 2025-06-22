package service

import (
	"errors"
	"log"
	"strings"

	"github.com/Hakeera/easy_erp/config"
	"github.com/Hakeera/easy_erp/model"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() ([]model.User, error) {
	log.Printf("Service GetUsers")

	var users []model.User


        db := config.GetDB()

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}


// Método principal que usa o banco configurado em config
func CreateUser(user *model.User) (*model.User, error) {
    log.Printf("Service CreateUser chamado para: %s", user.Username)
    
    // Obtém a instância do Banco de Dados
    db := config.GetDB()
    
    // Verificação se user não é nil
    if user == nil {
        log.Printf("ERRO: user é nil")
        return nil, errors.New("usuário não pode ser nil")
    }
    
    // Validações de negócio
    if err := validateUser(user); err != nil {
        log.Printf("Validação falhou: %v", err)
        return nil, err
    }
    
    // Verificar se usuário já existe 
    if userExists(user.Email, user.Username) {
        log.Printf("Usuário já existe: %s / %s", user.Email, user.Username)
        return nil, errors.New("usuário com este email ou username já existe")
    }

    if !isValidEmail(user.Email) {
	    log.Println("Validação de e-mail falhou")
	    return nil, errors.New("e-mail inválido")
    }

    // Hash da senha
    hashedPassword, err := hashPassword(user.Password)
    if err != nil {
        log.Printf("Erro ao fazer hash da senha: %v", err)
        return nil, errors.New("erro ao processar senha")
    }
    user.Password = hashedPassword
    
    // Criar usuário com GORM usando o banco configurado
    log.Printf("Tentando criar usuário no banco...")
    if err := db.Create(user).Error; err != nil {
        log.Printf("Erro no banco de dados: %v", err)
        return nil, err
    }
    
    log.Printf("Usuário criado no banco: ID=%d, Username=%s", user.ID, user.Username)
    return user, nil
}

// Função validateUser
func validateUser(user *model.User) error {
    if user == nil {
        return errors.New("usuário não pode ser nil")
    }
    
    if user.Username == "" {
        return errors.New("username é obrigatório")
    }
    
    if user.Email == "" {
        return errors.New("email é obrigatório")
    }
    
    if user.Password == "" {
        return errors.New("senha é obrigatória")
    }
    
    // Validações adicionais
    if len(user.Username) < 3 {
        return errors.New("username deve ter pelo menos 3 caracteres")
    }
    
    if len(user.Password) < 6 {
        return errors.New("senha deve ter pelo menos 6 caracteres")
    }
    
    return nil
}

// Função userExists usando o banco configurado
func userExists(email, username string) bool {
    db := config.GetDB()
    
    var count int64
   
    // Verificar email
    if err := db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
        log.Printf("Erro ao verificar email existente: %v", err)
        return false
    }
    
    if count > 0 {
        log.Printf("Email já existe: %s", email)
        return true
    }
    
    // Verificar username
    if err := db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
        log.Printf("Erro ao verificar username existente: %v", err)
        return false
    }
    
    if count > 0 {
        log.Printf("Username já existe: %s", username)
        return true
    }
    
    return false
}

// Função hashPassword
func hashPassword(password string) (string, error) {
    if password == "" {
        return "", errors.New("senha não pode estar vazia")
    }
    
    // Gerar hash da senha
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Erro ao gerar hash da senha: %v", err)
        return "", err
    }
    
    return string(hashedBytes), nil
}

// Validação de email
func isValidEmail(email string) bool {
	// Validação básica 
	return len(email) > 0 && 
	len(email) <= 254 && 
	strings.Contains(email, "@") && 
	strings.Contains(email, ".")
}
