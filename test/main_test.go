package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/task.manager.api/database"
	"github.com/hugaojanuario/task.manager.api/internal/handler/controllers"
	"github.com/hugaojanuario/task.manager.api/model"
	"github.com/stretchr/testify/assert"
)

func SetupTest() *gin.Engine {
	database.ConectingOnDatabase()
	routes := gin.Default()

	return routes
}

func TestVerifyFindTaskByIdStatusCode(t *testing.T) {
	r := SetupTest()
	r.GET("/task/:id", controllers.FindTaskById)

	req, _ := http.NewRequest("GET", "/task/7", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseTask model.Task
	err := json.Unmarshal(resp.Body.Bytes(), &responseTask)
	assert.NoError(t, err)

	assert.Equal(t, uint(7), responseTask.ID)
	assert.Equal(t, "fazer", responseTask.Title)
	assert.Equal(t, "trocar a luz", responseTask.Description)
	assert.Equal(t, model.StatusPedende, responseTask.Status)
}
