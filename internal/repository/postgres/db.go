package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// @ViitoJooj
// Migrations fica muito mais facil caso você precise trocar o banco de dados
func InitDatabase(db *sql.DB) {
	query, err := os.ReadFile("./migrations/0001_init_database.sql")
	if err != nil {
		log.Fatalf("Erro ao ler migration: %v", err)
	}

	_, err = db.Exec(string(query))
	if err != nil {
		log.Fatalf("Erro ao executar migration: %v", err)
	}

	log.Println("Banco inicializado com sucesso!")
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

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return nil
}
