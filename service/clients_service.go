package services

import (
	"errors"

	"github.com/Hakeera/easy_erp/configuration"
	models "github.com/Hakeera/easy_erp/model"
)

func CreateClient(client *models.Client) error {
	result := configuration.DB.Create(client)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetClients() ([]models.Client, error) {
	var clients []models.Client
	result := configuration.DB.Find(&clients)
	if result.Error != nil {
		return nil, result.Error
	}
	return clients, nil
}

func UpdateClient(id uint, updatedData *models.Client) error {
	var client models.Client
	result := configuration.DB.First(&client, id)
	if result.Error != nil {
		return errors.New("cliente não encontrado")
	}

	client.Name = updatedData.Name
	client.Email = updatedData.Email
	client.Phone = updatedData.Phone

	configuration.DB.Save(&client)
	return nil
}

func DeleteClient(id uint) error {
	result := configuration.DB.Delete(&models.Client{}, id)
	if result.Error != nil {
		return errors.New("erro ao excluir cliente")
	}
	return nil
}

