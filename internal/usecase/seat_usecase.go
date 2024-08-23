package usecase

import (
	"cinema/internal/model"
	"cinema/internal/repository"
)

type(
	SeatUsecase interface{
		AddSeat(seat model.SeatInput) error
		UpdateStatusSeat(seatRequest model.UpdateSeat) error
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

func (su seatUsecase) UpdateStatusSeat(seatRequest model.UpdateSeat) error{

	seat, err := su.seatRepository.FindSeatByID(seatRequest.ID)

	if err != nil{
		return err
	}

	seat.Status = model.SeatStatus(seatRequest.Status)

	err = su.seatRepository.UpdateStatusSeat(seat)

	if err != nil{
		return err
	}

	return nil
}