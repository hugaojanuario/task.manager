package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/task.manager.api/database"
	"github.com/hugaojanuario/task.manager.api/internal/handler/controllers"
	"github.com/stretchr/testify/assert"
)

func SetupTest() *gin.Engine {
	routes := gin.Default()

	return routes
}

func TestFindAllTasks(t *testing.T) {
	database.ConectingOnDatabase()
	r := SetupTest()
	r.GET("/task", controllers.FindTasks)

	req, _ := http.NewRequest("GET", "/task", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	fmt.Println(resp.Body)
}
