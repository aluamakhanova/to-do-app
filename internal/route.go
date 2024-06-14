package internal

import (
	"github.com/aluamakhanova/to-do-app/config"
	"github.com/aluamakhanova/to-do-app/internal/tasks"
	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	db := config.ConnectDB()

	taskController := tasks.NewTaskController(tasks.NewService(tasks.NewRepository(db)))

	route.POST("/tasks", taskController.CreateTask)
	route.GET("/tasks", taskController.GetAllTasks)
	route.PUT("/tasks/:taskId", taskController.UpdateTask)
	route.DELETE("/tasks/:taskId", taskController.DeleteTask)

	route.Run()
}
