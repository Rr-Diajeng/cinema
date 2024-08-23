package repository

import (
	"cinema/internal/model"

	"gorm.io/gorm"
)

type (
	SeatRepository interface {
		AddSeat(seat model.Seats) error
		UpdateStatusSeat(seat model.Seats) error
		FindSeatByID(id uint) (model.Seats, error)
		GetSeatByStatus(status string) ([]model.Seats, error)
		FindClassByID(id uint) (model.Class, error)
		FindCinemaStudiosByID(id uint) (model.CinemaStudios, error)
		GetSeatByClass(id uint)([]model.Seats, error)
	}

	seatRepository struct {
		db *gorm.DB
	}
)

func NewSeatRepository(db *gorm.DB) SeatRepository {
	return seatRepository{
		db: db,
	}
}

func (sr seatRepository) FindSeatByID(id uint) (model.Seats, error) {
	var seat model.Seats

	err := sr.db.Preload("Tickets").First(&seat, id).Error

	if err != nil {
		return seat, err
	}

	return seat, nil
}

func (sr seatRepository) FindClassByID(id uint) (model.Class, error){
	var class model.Class

	err := sr.db.Where("id = ?", id).First(&class).Error

	if err != nil{
		return class, err
	}

	return class, nil
}

func (sr seatRepository) FindCinemaStudiosByID(id uint) (model.CinemaStudios, error){
	var cinemaStudio model.CinemaStudios

	err := sr.db.First(&cinemaStudio, id).Error

	if err != nil{
		return cinemaStudio, err
	}

	return cinemaStudio, nil
}

func (sr seatRepository) AddSeat(seat model.Seats) error {

	err := sr.db.Save(&seat).Error

	if err != nil {
		return err
	}

	return nil
}

func (sr seatRepository) UpdateStatusSeat(seat model.Seats) error {
	err := sr.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&seat).Error

	if err != nil {
		return err
	}

	return nil
}

func (sr seatRepository) GetSeatByStatus(status string) ([]model.Seats, error) {
	var seats []model.Seats

	err := sr.db.Preload("Tickets").Where("status = ?", status).Find(&seats).Error

	if err != nil {
		return nil, err
	}

	return seats, nil
}

func (sr seatRepository) GetSeatByClass(id uint)([]model.Seats, error){
	var seats []model.Seats

	err := sr.db.Where("class_id = ?", id).Find(&seats).Error

	if err != nil{
		return nil, err
	}

	return seats, nil
}