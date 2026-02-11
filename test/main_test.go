package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/task.manager.api/database"
	"github.com/hugaojanuario/task.manager.api/internal/handler/controllers"
	"github.com/hugaojanuario/task.manager.api/model"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTest() *gin.Engine {
	routes := gin.Default()

	return routes
}

func CreatedMockTask() {
	task := model.Task{Title: "tasktest", Description: "this task is one test"}
	database.DB.Create(&task)
	ID = int(task.ID)
}

func DeleteMockTask() {
	var task model.Task
	database.DB.Delete(&task, ID)
}

func TestFindAllTasks(t *testing.T) {
	database.ConectingOnDatabase()
	CreatedMockTask()
	defer DeleteMockTask()
	r := SetupTest()
	r.GET("/task", controllers.FindTasks)

	req, _ := http.NewRequest("GET", "/task", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	fmt.Println(resp.Body)
}

func TestFindTaskByIdHandler(t *testing.T) {
	database.ConectingOnDatabase()
	CreatedMockTask()
	defer DeleteMockTask()
	r := SetupTest()
	r.GET("/task/:id", controllers.FindTaskById)
	path := "/task/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var taskMock model.Task
	json.Unmarshal(resp.Body.Bytes(), &taskMock)
	assert.Equal(t, "tasktest", taskMock.Title)
	assert.Equal(t, "this task is one test", taskMock.Description)
}
