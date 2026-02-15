package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/task.manager.api/internal/handler/controllers"
)

func HandlerRequest() {
	r := gin.Default()
	r.POST("/task", controllers.CreatedTask)
	r.GET("/task", controllers.FindAllTasks)
	r.GET("/task/:id", controllers.FindTaskById)
	r.PUT("/task/:id", controllers.PutTaskById)
	r.DELETE("task/:id", controllers.DeleteById)
	r.Run()
}
