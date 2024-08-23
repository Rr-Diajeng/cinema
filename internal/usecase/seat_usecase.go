package usecase

import (
	"cinema/internal/model"
	"cinema/internal/repository"
)

type(
	SeatUsecase interface{
		AddSeat(seat model.SeatInput) error
	}

	seatUsecase struct{
		seatRepository repository.SeatRepository
	}
)

func NewSeatUsecase(seatRepository repository.SeatRepository) SeatUsecase{
	return seatUsecase{
		seatRepository: seatRepository,
	}
}

func (su seatUsecase) AddSeat(seat model.SeatInput) error{
	err := su.seatRepository.AddSeat(model.Seats{
		CinemaStudiosID: seat.CinemaStudiosID,
		ClassID: seat.ClassID,
		SeatNumber: seat.SeatNumber,
		Status: model.SeatStatus(seat.Status),
	})

	if err != nil{
		return err
	}

	return nil
}