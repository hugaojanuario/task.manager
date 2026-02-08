package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/task.manager.api/database"
	"github.com/hugaojanuario/task.manager.api/model"
)

func CreatedTask(c *gin.Context){
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err,
		})
		return
	}
	database.DB.Create(&task)
	task.Status = model.StatusPedende
	c.JSON(http.StatusCreated, task)
}

func FindTasks(c *gin.Context){
	var tasks []model.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func FindTaskById(c *gin.Context){
	var task model.Task
	id := c.Params.ByName("id")
	database.DB.First(&task, id)

	if task.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "task not found",
		})
		return
	}
	c.JSON(http.StatusOK, task)

}

func PutTaskById(c *gin.Context){
	var task model.Task
	id := c.Params.ByName("id")
	database.DB.First(&task, id)

	if err := c.ShouldBindJSON(&task); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err,
		})
		return
	}
	database.DB.Model(&task).UpdateColumns(task)
	c.JSON(http.StatusOK, task)
}

func DeleteById(c *gin.Context){
	var task model.Task
	id := c.Params.ByName("id")
	

	database.DB.First(&task, id)
	database.DB.Delete(&task)
	c.JSON(http.StatusOK, task)
}