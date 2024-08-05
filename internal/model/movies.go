package model

import (
	"time"

	"gorm.io/gorm"
)

type Movies struct {
	gorm.Model
	ID          uint          `gorm:"type:bigint;primary_key;auto_increment"`
	Title       string        `gorm:"type:varchar;not_null"`
	Genres      []Genres      `gorm:"many2many:genre_movies;"`
	Duration    time.Duration `gorm:"type:bigint"`
	ReleaseDate time.Time     `gorm:"type:timestamp;not_null"`
	Synopsis    string        `gorm:"type:text"`
	BasePrice   float32       `gorm:"type:float"`
	Tickets     []Tickets     `gorm:"foreignKey:MovieID"`
}