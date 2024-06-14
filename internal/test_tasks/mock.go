package test_tasks

import (
	"github.com/aluamakhanova/to-do-app/internal/tasks"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(task *tasks.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockRepository) FindAll() ([]tasks.Task, error) {
	args := m.Called()
	return args.Get(0).([]tasks.Task), args.Error(1)
}

func (m *MockRepository) FindByID(id uint) (tasks.Task, error) {
	args := m.Called(id)
	return args.Get(0).(tasks.Task), args.Error(1)
}

func (m *MockRepository) Update(task *tasks.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
