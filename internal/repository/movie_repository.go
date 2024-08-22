package repository

import (
	"cinema/internal/model"

	"gorm.io/gorm"
)

type (
	MovieRepository interface {
		AddMovie(movie model.Movies) (err error)
		FindGenreByID(id uint, genre *model.Genres) error
		FindMovieByID(id uint) (model.Movies, error)
		UpdateMovie(movie model.Movies) error
		ClearGenres(id uint) error
	}

	movieRepository struct {
		db *gorm.DB
	}
)

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return movieRepository{
		db: db,
	}
}

func (mr movieRepository) AddMovie(movie model.Movies) (err error) {

	if err := mr.db.Save(&movie).Error; err != nil {
		return err
	}

	return nil
}

func (mr movieRepository) FindGenreByID(id uint, genre *model.Genres) error {
	return mr.db.Where("id = ?", id).First(&genre).Error
}

func (mr movieRepository) FindMovieByID(id uint) (model.Movies, error) {
	var movie model.Movies

	err := mr.db.Preload("Genres").First(&movie, id).Error

	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (mr movieRepository) UpdateMovie(movie model.Movies) error {
	return mr.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&movie).Error
}

func (mr movieRepository) ClearGenres(id uint) error {
	return mr.db.Model(&model.Movies{ID: id}).Association("Genres").Clear()
}