package service

import (
	"errors"
	"fmt"

	"github.com/Hakeera/easy_erp/config"
	"github.com/Hakeera/easy_erp/model"
	"gorm.io/gorm"
)

// DB é a instância de conexão com o banco de dados
var DB *gorm.DB

// CreateClient adiciona um novo cliente ao banco de dados usando GORM
func CreateClient(client *model.Client) error {
	db := config.GetDB()
	
	// Cria o cliente no banco
	result := db.Create(&client)

	// Verifica erro de duplicidade (constraint unique)
	if result.Error != nil {
		if result.Error.Error() == 
			`ERROR: duplicate key value violates unique constraint "uk_client_document" (SQLSTATE 23505)` {
			return errors.New("já existe um cliente com este documento")
		}
		return fmt.Errorf("erro ao salvar no banco: %w", result.Error)
	}

	return nil
}

// GetClients retorna todos os clientes do banco de dados
func GetClients() ([]model.Client, error) {
	var clients []model.Client
	
	// Instancia o Banco de Dados
	db := config.GetDB()
	
	// O GORM vai retornar todos os clientes ordenados por nome
	result := db.Find(&clients)
	if result.Error != nil {
		fmt.Println("erro ao Buscar clientes: %w", result.Error)
		return nil, result.Error
	}
	
	return clients, nil
}

// GetClientByID busca um cliente pelo ID
func GetClientByID(id uint) (*model.Client, error) {
	var client model.Client
	
	result := DB.First(&client, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("cliente não encontrado")
		}
		return nil, result.Error
	}
	
	return &client, nil
}

func UpdateClient(id uint, updatedClient *model.Client) error {
	// Atualiza com base no ID informado
	return DB.Model(&model.Client{}).Where("id = ?", id).Updates(updatedClient).Error
}


// DeleteClient remove um cliente do banco de dados (soft delete com GORM)
func DeleteClient(id uint) error {
	result := DB.Delete(&model.Client{}, id)
	return result.Error
}
