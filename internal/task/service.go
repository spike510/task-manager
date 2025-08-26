package task

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateTask(userID int, title, description string) (*Task, error) {
	return s.repo.Create(userID, title, description)
}

func (s *Service) GetTasks(userID int) ([]Task, error) {
	return s.repo.GetAllByUser(userID)
}

func (s *Service) UpdateTask(id int, userID int, title, description string, done bool) (*Task, error) {
	return s.repo.Update(id, userID, title, description, done)
}

func (s *Service) DeleteTask(id int, userID int) error {
	return s.repo.Delete(id, userID)
}
