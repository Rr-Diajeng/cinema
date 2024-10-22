package model

import "gorm.io/gorm"

type SeatStatus string

const (
	Available SeatStatus = "available"
	Booked    SeatStatus = "booked"
)

type Seats struct {
	gorm.Model
	ID              uint `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	CinemaStudiosID uint
	ClassID         uint
	SeatNumber      string        `gorm:"type:varchar;not_null"`
	Status          SeatStatus    `gorm:"type:seats_status;not_null;default:'available'"`
	CinemaStudios   CinemaStudios `gorm:"foreignKey:CinemaStudiosID"`
	Class           Class         `gorm:"foreignKey:ClassID"`
	Tickets         []Tickets     `gorm:"foreignKey:SeatID"`
}

type SeatInput struct {
	CinemaStudiosID uint   `json:"cinemaStudiosID" binding:"required"`
	ClassID         uint   `json:"classID" binding:"required"`
	SeatNumber      string `json:"seatNumber" binding:"required"`
	Status          string `json:"status" binding:"required,oneof='available' 'booked'"`
}

type UpdateSeat struct {
	ID     uint   `json:"id" binding:"required"`
	Status string `json:"status" binding:"required,oneof='available' 'booked'"`
}

type SeatRequestByStatus struct {
	Status string `json:"status" binding:"required,oneof='available' 'booked'"`
}

type SeatResponse struct {
	CinemaStudios string `json:"cinemaStudios"`
	Class         string `json:"class"`
	SeatNumber    string `json:"seatNumber"`
	Status        string `json:"status"`
}

type SeatRequestByClass struct {
	ClassID uint `json:"class_id"`
}

type SeatRequestByCinemaStudios struct {
	CinemaStudiosID uint `json:"cinema_studios_id"`
}

type IDSeatRequest struct{
	ID uint `json:"id"`
}