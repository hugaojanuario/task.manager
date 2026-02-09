package main

import (
	"io/ioutil"
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

	req, _ := http.NewRequest("GET", "/task/3", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "error menssage: 'deveriam ser iguais'")
	mockResponse, _ := ioutil.ReadAll(resp.Body)
	responseBody := &model.Task{}
	assert.Equal(t, mockResponse, responseBody)
}
