package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/task.manager.api/database"
	"github.com/hugaojanuario/task.manager.api/internal/domain"
	"github.com/hugaojanuario/task.manager.api/internal/handler"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTest() *gin.Engine {
	routes := gin.Default()

	return routes
}

func CreatedMockTask() {
	task := domain.Task{Title: "tasktest", Description: "this task is one test"}
	database.DB.Create(&task)
	ID = int(task.ID)
}

func DeleteMockTask() {
	var task domain.Task
	database.DB.Delete(&task, ID)
}

func TestFindAllTasks(t *testing.T) {
	database.ConectingOnDatabase()
	CreatedMockTask()
	defer DeleteMockTask()
	r := SetupTest()
	r.GET("/task", handler.FindAllTasks)
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
	r.GET("/task/:id", handler.FindTaskById)
	path := "/task/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var taskMock domain.Task
	json.Unmarshal(resp.Body.Bytes(), &taskMock)
	assert.Equal(t, "tasktest", taskMock.Title)
	assert.Equal(t, "this task is one test", taskMock.Description)
}

func TestDeleteByIdHandler(t *testing.T) {
	database.ConectingOnDatabase()
	CreatedMockTask()
	DeleteMockTask()
	r := SetupTest()
	r.DELETE("/task/:id", handler.DeleteById)
	path := "/task/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestPutByIdHandler(t *testing.T) {
	database.ConectingOnDatabase()
	CreatedMockTask()
	defer DeleteMockTask()
	r := SetupTest()
	r.PUT("/task/:id", handler.PutTaskById)
	task := domain.Task{Title: "tasktest", Description: "this task is one test2"}
	taskInJson, _ := json.Marshal(task)
	path := "/task/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(taskInJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var taskMock domain.Task
	json.Unmarshal(resp.Body.Bytes(), &taskMock)
	assert.Equal(t, "tasktest", taskMock.Title)
	assert.Equal(t, "this task is one test2", taskMock.Description)
}
