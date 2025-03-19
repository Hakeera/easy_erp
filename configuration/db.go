package config

import (
	"database/sql"
	"fmt"
	"log"
)

// Configuração da conexão com o banco
const (
	host     = "localhost"
	port     = 5432  
	user     = "user"
	password = "password"
	dbname   = "crud"
)

// Conectar ao banco
func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco:", err)
	}

	// Verificar se a conexão está ativa
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	fmt.Println("Banco de dados conectado com sucesso!")
	return db
}

