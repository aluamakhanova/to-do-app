package tasks

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(task *Task) error
	FindAll() ([]Task, error)
	FindByID(id uint) (Task, error)
	Update(task *Task) error
	Delete(id uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (repo *TaskRepository) Create(task *Task) error {
	return repo.db.Create(task).Error
}

func (repo *TaskRepository) FindAll() ([]Task, error) {
	var tasks []Task
	result := repo.db.Find(&tasks)
	return tasks, result.Error
}

func (repo *TaskRepository) FindByID(id uint) (Task, error) {
	var task Task
	result := repo.db.Where("id = ?", id).First(&task)
	return task, result.Error
}

func (repo *TaskRepository) Update(task *Task) error {
	return repo.db.Save(task).Error
}

func (repo *TaskRepository) Delete(id uint) error {
	var task Task
	return repo.db.Where("id = ?", id).Delete(&task).Error
}
