package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func CreateTable() {
	query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        status VARCHAR(50)
    );`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	log.Println("Tabela 'tasks' verificada/criada com sucesso!")
}

func ConectingOnDatabase() error {
	connStr := "host=localhost user=admin password=123456 dbname=taskmanager port=5432 sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("erro ao abrir driver: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("erro ao conectar no banco (ping): %v", err)
	}

	DB.SetMaxOpenConns(10)           // Máximo de conexões abertas simultaneamente
	DB.SetMaxIdleConns(5)            // Máximo de conexões paradas esperando uso
	DB.SetConnMaxLifetime(time.Hour) // Tempo máximo que uma conexão pode viver

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return nil
}
