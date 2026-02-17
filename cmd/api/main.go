package main

import (
	"github.com/hugaojanuario/task.manager.api/internal/http"
	"github.com/hugaojanuario/task.manager.api/internal/repository/postgres"
)

// @ViitoJooj
func main() {
	// Abre conexão com o banco
	postgres.ConectingOnDatabase()
	// Executa migrations usando a conexão aberta
	postgres.InitDatabase(postgres.DB)
	// Inicia as rotas
	http.HandlerRequest()
}
