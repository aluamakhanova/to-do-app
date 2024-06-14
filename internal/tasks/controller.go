package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

type TaskController struct {
	service *Service
}

func NewTaskController(service *Service) *TaskController {
	return &TaskController{service: service}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var data TaskRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.service.CreateTask(data.Title, data.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := cast.ToUint(c.Param("taskId"))
	var data TaskRequest
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.service.UpdateTask(id, data.Title, data.Description, data.Completed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := cast.ToUint(c.Param("taskId"))
	if err := tc.service.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Success", "data": id})
}
