package model

import "gorm.io/gorm"

type Tickets struct {
	gorm.Model
	ID              uint `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	UserID          uint
	MovieID         uint
	CinemaStudiosID uint
	SeatID          uint
	TransactionID   uint
	Price           float32       `gorm:"type:float"`
	Barcode         string        `gorm:"type:varchar;not_null"`
	Users           Users         `gorm:"foreignKey:UserID"`
	Movies          Movies        `gorm:"foreignKey:MovieID"`
	CinemaStudios   CinemaStudios `gorm:"foreignKey:CinemaStudiosID"`
	Seats           Seats         `gorm:"foreignKey:SeatID"`
	Transactions    Transactions  `gorm:"foreignKey:TransactionID"`
}