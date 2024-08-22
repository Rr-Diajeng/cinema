package repository

import (
	"cinema/internal/model"
	"time"

	"gorm.io/gorm"
)

type (
	MovieRepository interface {
		AddMovie(movie model.Movies) (err error)
		FindGenreByID(id uint, genre *model.Genres) error
		FindMovieByID(id uint) (model.Movies, error)
		UpdateMovie(movie model.Movies) error
		ClearGenres(id uint) error
		GetMovieInSchedule(exactTime time.Time) ([]model.Movies, error)
		DeleteMovieByID(id uint) error
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

func (mr movieRepository) GetMovieInSchedule(exactTime time.Time) ([]model.Movies, error) {
	var movies []model.Movies
	err := mr.db.Preload("Genres").Preload("Tickets").Where("start_date <= ? AND end_date >= ?", exactTime, exactTime).Find(&movies).Error
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (mr movieRepository) DeleteMovieByID(id uint) error{
	err := mr.db.Delete(&model.Movies{}, id).Error

	if err != nil{
		return err
	}

	return nil
}