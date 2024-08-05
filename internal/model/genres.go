package model

type Genres struct {
	ID     uint     `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	Type   string   `gorm:"type:string;not_null"`
	Movies []Movies `gorm:"many2many:genre_movies;"`
}