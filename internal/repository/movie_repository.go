package repository

import (
	"cinema/internal/model"
	"context"

	"gorm.io/gorm"
)

type (
	MovieRepository interface{
		AddMovie(c context.Context, movie model.Movies) (err error)
		FindGenreByID(c context.Context, id uint, genre *model.Genres) error
	}

	movieRepository struct{
		db *gorm.DB
	}
)

func NewMovieRepository(db *gorm.DB) MovieRepository{
	return movieRepository{
		db: db,
	}
}

func (mr movieRepository) AddMovie(c context.Context, movie model.Movies) (err error){

	if err := mr.db.Save(&movie).Error; err != nil{
		return err
	}

	return nil
}

func (mr movieRepository) FindGenreByID(c context.Context, id uint, genre *model.Genres) error{
    return mr.db.Where("id = ?", id).First(genre).Error
}