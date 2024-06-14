package tasks

type Service struct {
	repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) CreateTask(title, description string) (*Task, error) {
	task := Task{Title: title, Description: description}
	err := s.repository.Create(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *Service) GetAllTasks() ([]Task, error) {
	return s.repository.FindAll()
}

func (s *Service) UpdateTask(id uint, title, description string, completed bool) (*Task, error) {
	task, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	task.Title = title
	task.Description = description
	task.Completed = completed
	err = s.repository.Update(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *Service) DeleteTask(id uint) error {
	return s.repository.Delete(id)
}
