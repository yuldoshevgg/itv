package movie

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(movie *Movie) error
	GetAll() ([]Movie, error)
	GetByID(id uint) (*Movie, error)
	Update(movie *Movie) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(movie *Movie) error {
	return r.db.Create(movie).Error
}

func (r *repository) GetAll() ([]Movie, error) {
	var movies []Movie
	err := r.db.Find(&movies).Error
	return movies, err
}

func (r *repository) GetByID(id uint) (*Movie, error) {
	var movie Movie
	err := r.db.First(&movie, id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *repository) Update(movie *Movie) error {
	return r.db.Save(movie).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Movie{}, id).Error
}
