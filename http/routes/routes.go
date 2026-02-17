package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/task.manager.api/internal/handler"
)

func HandlerRequest() {
	r := gin.Default()
	r.POST("/task", handler.CreatedTask)
	r.GET("/task", handler.FindAllTasks)
	r.GET("/task/:id", handler.FindTaskById)
	r.PUT("/task/:id", handler.PutTaskById)
	r.DELETE("task/:id", handler.DeleteById)
	r.Run()
}
