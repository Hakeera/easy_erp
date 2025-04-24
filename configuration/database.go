package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Hakeera/easy_erp/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// InitDB inicializa a conexão com o banco de dados
func InitDB() {
	once.Do(func() {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPassword, dbName, dbPort,
		)

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		}

		// Migrar a tabela usando o próprio DB
		if err := DB.AutoMigrate(&model.FichaModeloDB{}); err != nil {
			log.Fatalf("Erro ao migrar banco: %v", err)
		}
	})
}

// GetDB retorna a instância do banco de dados inicializada
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Banco de dados não foi inicializado!")
	}
	return DB
}

