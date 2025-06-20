package service

import (
	"errors"
	"log"
	"strings"

	"github.com/Hakeera/easy_erp/model"
	"golang.org/x/crypto/bcrypt"
)

// Método principal que seu controller chama
func CreateUser(user *model.User) (*model.User, error) {
	log.Printf("Service CreateUser chamado para: %s", user.Username)

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

	// Hash da senha
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		log.Printf("Erro ao fazer hash da senha: %v", err)
		return nil, errors.New("erro ao processar senha")
	}
	user.Password = hashedPassword

	// Criar usuário com GORM
	if err := DB.Create(user).Error; err != nil {
		log.Printf("Erro no banco de dados: %v", err)
		return nil, err
	}

	log.Printf("Usuário criado no banco: ID=%d, Username=%s", user.ID, user.Username)
	return user, nil
}

// Verificar se usuário existe
func userExists(email, username string) bool {
	var count int64
	DB.Model(&model.User{}).Where("email = ? OR username = ?", email, username).Count(&count)
	return count > 0
}

// Validações de negócio
func validateUser(user *model.User) error {
	if len(user.Username) < 3 {
		return errors.New("username deve ter pelo menos 3 caracteres")
	}

	if len(user.Password) < 6 {
		return errors.New("senha deve ter pelo menos 6 caracteres")
	}

	// Validação básica de email
	if !isValidEmail(user.Email) {
		return errors.New("formato de email inválido")
	}

	return nil
}

// Hash da senha com bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Validação de email
func isValidEmail(email string) bool {
	// Validação básica - você pode usar uma lib mais robusta
	return len(email) > 0 && 
	len(email) <= 254 && 
	strings.Contains(email, "@") && 
	strings.Contains(email, ".")
}
