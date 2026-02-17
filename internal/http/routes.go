// @ViitoJooj
// rotas dentro de http

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/task.manager.api/internal/http/handler"
	"github.com/hugaojanuario/task.manager.api/internal/repository"
	"github.com/hugaojanuario/task.manager.api/internal/repository/postgres"
	"github.com/hugaojanuario/task.manager.api/internal/service"
)

func HandlerRequest() {
	r := gin.Default()
	repo := repository.NewTaskRepository(postgres.DB)
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	r.POST("/tasks", handler.Create)
	r.GET("/tasks", handler.FindAll)
	r.GET("/tasks/:id", handler.FindByID)
	r.PUT("/tasks/:id", handler.Update)
	r.DELETE("/tasks/:id", handler.Delete)
	r.Run()
}
