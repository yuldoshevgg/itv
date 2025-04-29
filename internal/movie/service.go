package movie

type Service interface {
	Create(movie *Movie) error
	GetAll() ([]Movie, error)
	GetByID(id uint) (*Movie, error)
	Update(id uint, updatedMovie *Movie) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(movie *Movie) error {
	return s.repo.Create(movie)
}

func (s *service) GetAll() ([]Movie, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id uint) (*Movie, error) {
	return s.repo.GetByID(id)
}

func (s *service) Update(id uint, updatedMovie *Movie) error {
	movie, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	movie.Title = updatedMovie.Title
	movie.Director = updatedMovie.Director
	movie.Year = updatedMovie.Year
	movie.Plot = updatedMovie.Plot

	return s.repo.Update(movie)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
