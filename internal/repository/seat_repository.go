package repository

import (
	"cinema/internal/model"

	"gorm.io/gorm"
)

type(
	SeatRepository interface{
		AddSeat(seat model.Seats) error
		UpdateStatusSeat(seat model.Seats) error
		FindSeatByID(id uint) (model.Seats, error)
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

func (sr seatRepository) FindSeatByID(id uint) (model.Seats, error){
	var seat model.Seats

	err := sr.db.Preload("Tickets").First(&seat, id).Error

	if err != nil{
		return seat, err
	}

	return seat, nil
}

func (sr seatRepository) AddSeat(seat model.Seats) error{

	err := sr.db.Save(&seat).Error

	if err != nil{
		return err
	}

	return nil
}

func (sr seatRepository) UpdateStatusSeat(seat model.Seats) error{
	err := sr.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&seat).Error

	if err != nil{
		return err
	}

	return nil
}