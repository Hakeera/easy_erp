package service

import (
	"encoding/json"
	"fmt"

	"github.com/Hakeera/easy_erp/config"
	"github.com/Hakeera/easy_erp/model"
)

func SalvarFichaModelo(ficha model.FichaModelo) error {
	db := config.GetDB()

	// Converte a ficha para JSON
	jsonBytes, err := json.Marshal(ficha)
	if err != nil {
		return fmt.Errorf("erro ao converter ficha para JSON: %w", err)
	}

	// Cria a struct para salvar no banco
	fichaDB := model.FichaModeloDB{
		Dados: jsonBytes,
	}

	// Salva no banco
	if err := db.Create(&fichaDB).Error; err != nil {
		return fmt.Errorf("erro ao salvar no banco: %w", err)
	}

	return nil
}

