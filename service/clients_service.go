package service

import (
	"errors"

	"github.com/Hakeera/easy_erp/configuration"
	"github.com/Hakeera/easy_erp/model"
)

// Criar Cliente
func CreateClient(client *model.Cliente) error {
	result := configuration.DB.Create(client) // Isso deve adicionar o cliente ao banco
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Buscar todos os clientes
func GetClients() ([]model.Cliente, error) {
	db := configuration.GetDB()
	var clients []model.Cliente
	if err := db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

// Atualizar Cliente
func UpdateClient(id uint, updatedData *model.Cliente) error {
	db := configuration.GetDB()
	var client model.Cliente

	if err := db.First(&client, id).Error; err != nil {
		return errors.New("cliente não encontrado")
	}

	// Atualizando os dados
	client.Nome = updatedData.Nome
	client.Email = updatedData.Email
	client.Telefone = updatedData.Telefone

	if err := db.Save(&client).Error; err != nil {
		return err
	}
	return nil
}

// Deletar Cliente
func DeleteClient(id uint) error {
	db := configuration.GetDB()
	if err := db.Delete(&model.Cliente{}, id).Error; err != nil {
		return errors.New("erro ao excluir cliente")
	}
	return nil
}

