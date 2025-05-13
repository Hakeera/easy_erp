package model

import (
	"time"

	"gorm.io/gorm"
)

// Client representa um cliente no sistema
type Client struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"type:varchar(255);not null"`
	TradingName    string         `json:"trading_name,omitempty" gorm:"type:varchar(255)"`
	Email          string         `json:"email,omitempty" gorm:"type:varchar(255)"`
	Phone          string         `json:"phone,omitempty" gorm:"type:varchar(20)"`
	Address        string         `json:"address,omitempty" gorm:"type:text"`
	DocumentType   string         `json:"document_type" gorm:"type:varchar(10);not null"`
	DocumentNumber string         `json:"document_number" gorm:"type:varchar(20);not null"`
	CEP            string         `json:"cep,omitempty" gorm:"type:char(8)"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at,omitempty"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName especifica o nome da tabela no banco de dados
func (Client) TableName() string {
	return "clients"
}

// BeforeCreate é um hook do GORM executado antes de criar um registro
func (c *Client) BeforeCreate(tx *gorm.DB) error {
	// Você pode adicionar validações ou transformações aqui
	return nil
}

// FormatDocument retorna o documento formatado de acordo com o tipo
func (c *Client) FormatDocument() string {
	if c.DocumentType == "CPF" {
		if len(c.DocumentNumber) == 11 {
			return c.DocumentNumber[:3] + "." + c.DocumentNumber[3:6] + "." + 
			       c.DocumentNumber[6:9] + "-" + c.DocumentNumber[9:]
		}
	} else if c.DocumentType == "CNPJ" {
		if len(c.DocumentNumber) == 14 {
			return c.DocumentNumber[:2] + "." + c.DocumentNumber[2:5] + "." + 
			       c.DocumentNumber[5:8] + "/" + c.DocumentNumber[8:12] + "-" + c.DocumentNumber[12:]
		}
	}
	return c.DocumentNumber
}

// FormatCEP retorna o CEP formatado (00000-000)
func (c *Client) FormatCEP() string {
	if len(c.CEP) == 8 {
		return c.CEP[:5] + "-" + c.CEP[5:]
	}
	return c.CEP
}
