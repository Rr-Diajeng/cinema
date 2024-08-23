package usecase

import (
	"cinema/internal/model"
	"cinema/internal/repository"
)

type(
	SeatUsecase interface{
		AddSeat(seat model.SeatInput) error
		UpdateStatusSeat(seatRequest model.UpdateSeat) error
		FindSeatByStatus(statusRequest model.SeatRequestByStatus) ([]model.SeatResponse, error)
		FindSeatByClass(classRequest model.SeatRequestByClass) ([]model.SeatResponse, error)
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

func (su seatUsecase) FindSeatByStatus(statusRequest model.SeatRequestByStatus) ([]model.SeatResponse, error){

	seats, err := su.seatRepository.GetSeatByStatus(statusRequest.Status)

	if err != nil{
		return nil, err
	}

	var seatResponse []model.SeatResponse

	for _, seat := range seats{
		cinemaStudio, err := su.seatRepository.FindCinemaStudiosByID(seat.CinemaStudiosID)
		if err != nil{
			return nil, err
		}

		class, err := su.seatRepository.FindClassByID(seat.ClassID)
		if err != nil{
			return nil, err
		}

		seatResponse = append(seatResponse, model.SeatResponse{
			CinemaStudios: cinemaStudio.Name,
			Class: class.Name,
            SeatNumber: seat.SeatNumber,
            Status: string(seat.Status),
		})
	}

	return seatResponse, nil
}

func (su seatUsecase) FindSeatByClass(classRequest model.SeatRequestByClass) ([]model.SeatResponse, error){
	seats, err := su.seatRepository.GetSeatByClass(classRequest.ClassID)

	if err != nil{
		return nil, err
	}

	var seatResponse []model.SeatResponse
	for _, seat := range seats{
		cinemaStudios, err := su.seatRepository.FindCinemaStudiosByID(seat.CinemaStudiosID)
		if err != nil{
			return nil, err
		}

		class, err := su.seatRepository.FindClassByID(seat.ClassID)
		if err != nil{
			return nil, err
		}

		seatResponse = append(seatResponse, model.SeatResponse{
			CinemaStudios: cinemaStudios.Name,
            Class: class.Name,
            SeatNumber: seat.SeatNumber,
            Status: string(seat.Status),
		})

	}

	return seatResponse, nil
}