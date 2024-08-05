package model

type CinemaStudios struct {
	ID      uint      `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	Name    string    `gorm:"type:varchar;not_null"`
	Tickets []Tickets `gorm:"foreignKey:CinemaStudiosID"`
	Seats   []Seats   `gorm:"foreignKey:CinemaStudiosID"`
}