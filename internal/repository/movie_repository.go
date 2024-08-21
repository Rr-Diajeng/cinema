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
		FindMovieByID(c context.Context, id uint) (model.Movies, error)
		UpdateMovie(c context.Context, movie model.Movies) error
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

func (mr movieRepository) FindMovieByID(c context.Context, id uint)(model.Movies, error){
	var movie model.Movies

	err := mr.db.Preload("Genres").First(&movie, id).Error

	if err != nil{
		return movie, err
	}

	return movie, nil
}

func (mr movieRepository) UpdateMovie(c context.Context, movie model.Movies) error{
	return mr.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&movie).Error
}