package test_tasks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aluamakhanova/to-do-app/internal/tasks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockRepo := new(MockRepository)
	service := tasks.NewService(mockRepo)
	controller := tasks.NewTaskController(service)

	mockRepo.On("Create", mock.Anything).Return(nil).Once()

	router.POST("/tasks", controller.CreateTask)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(gin.H{"title": "Test Task", "description": "Test Description"})
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Test Task")

	// Verify that expectations were met
	mockRepo.AssertExpectations(t)
}
