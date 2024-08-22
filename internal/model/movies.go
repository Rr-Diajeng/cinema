package model

import (
	"time"

	"gorm.io/gorm"
)

type Movies struct {
	gorm.Model
	ID          uint      `gorm:"type:bigint;primary_key;auto_increment"`
	Title       string    `gorm:"type:varchar;not_null"`
	Genres      []Genres  `gorm:"many2many:genre_movies;"`
	Duration    int64     `gorm:"type:bigint"`
	ReleaseDate time.Time `gorm:"type:timestamp;not_null"`
	Synopsis    string    `gorm:"type:text"`
	BasePrice   float32   `gorm:"type:float"`
	Tickets     []Tickets `gorm:"foreignKey:MovieID"`
	StartDate   time.Time `gorm:"type:timestamp;not_null"`
	EndDate     time.Time `gorm:"type:timestamp;not_null"`
}

type AddMovieRequest struct {
	Title       string    `json:"title"`
	Genres      []uint    `json:"genres"`
	Duration    string    `json:"duration"`
	ReleaseDate time.Time `json:"releaseDate"`
	Synopsis    string    `json:"synopsis"`
	BasePrice   float32   `json:"basePrice"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

type UpdateMovieRequest struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Genres      []uint    `json:"genres"`
	Duration    string    `json:"duration"`
	ReleaseDate time.Time `json:"releaseDate"`
	Synopsis    string    `json:"synopsis"`
	BasePrice   float32   `json:"basePrice"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

type MovieResponse struct {
	Title       string    `json:"title"`
	Genres      []Genres  `json:"genres"`
	Duration    int64     `json:"duration"`
	ReleaseDate time.Time `json:"releaseDate"`
	Synopsis    string    `json:"synopsis"`
	BasePrice   float32   `json:"basePrice"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

type OneMovieRequest struct {
	ID uint `json:"id"`
}

type MovieInScheduleRequest struct {
	ExactTime time.Time `json:"time"`
}
