package repository

import (
	"cinema/internal/model"

	"gorm.io/gorm"
)

type(
	SeatRepository interface{
		AddSeat(seat model.Seats) error
	}

	seatRepository struct{
		db *gorm.DB
	}
)

func NewSeatRepository (db *gorm.DB) SeatRepository{
	return seatRepository{
		db: db,
	}
}

func (sr seatRepository) AddSeat(seat model.Seats) error{

	err := sr.db.Save(&seat).Error

	if err != nil{
		return err
	}

	return nil
}