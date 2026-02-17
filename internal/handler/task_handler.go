package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/task.manager.api/database"
	"github.com/hugaojanuario/task.manager.api/internal/domain"
	"github.com/hugaojanuario/task.manager.api/internal/service"
)

func CreatedTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := service.ValidationTask(&task); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	task.Status = domain.StatusPedende
	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id`
	var id int64
	err := database.DB.QueryRow(query, task.Title, task.Description, task.Status).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco: " + err.Error()})
		return
	}
	task.ID = uint(id)
	c.JSON(http.StatusCreated, task)
}

func FindAllTasks(c *gin.Context) {
	// 1. Definimos a Query
	// DICA DE OURO: Evite "SELECT *". Liste os campos explicitamente.
	// Isso garante que a ordem do banco bata com a ordem do seu Scan.
	query := `SELECT id, title, description, status FROM tasks`

	// 2. Executamos a Query (Note que é .Query e não .QueryRow)
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefas: " + err.Error()})
		return
	}
	// IMPORTANTE: Fecha a conexão quando a função terminar para não vazar memória
	defer rows.Close()

	// 3. Inicializamos a lista (slice) onde vamos guardar os resultados
	// make([]Tipo, 0) cria uma lista vazia, pronta para uso
	tasks := make([]domain.Task, 0)

	// 4. O Loop: "Enquanto tiver uma próxima linha..."
	for rows.Next() {
		var task domain.Task

		// 5. O Scan: Mapeia as colunas do banco para os campos da Struct
		// A ORDEM AQUI TEM QUE SER IGUAL AO SELECT LÁ EM CIMA!
		// 1º id, 2º title, 3º description, 4º status
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler dados: " + err.Error()})
			return
		}

		// Adiciona a tarefa lida na lista final
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao iterar linhas: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func FindTaskById(c *gin.Context) {
	var task domain.Task
	id := c.Params.ByName("id")

	query := `SELECT id, title, description, status FROM tasks WHERE id = $1`
	err := database.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Task not found", // 404
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar no banco: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func PutTaskById(c *gin.Context) {
	var task domain.Task
	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValidationTask(&task); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE tasks SET title=$1, description=$2, status=$3 WHERE id=$4`

	result, err := database.DB.Exec(query, task.Title, task.Description, task.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar: " + err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func DeleteById(c *gin.Context) {
	id := c.Params.ByName("id")

	query := `DELETE FROM tasks WHERE id=$1`

	result, err := database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar: " + err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
		"id":      id,
	})
}
