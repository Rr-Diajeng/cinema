package main

import (
	"cinema/internal/database"
	"cinema/internal/handler"
	"cinema/internal/repository"
	"cinema/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main(){
    r := gin.Default()

    db := database.GetDBInstance()

    userRepository := repository.NewUserRepository(db)
    userUsecase := usecase.NewUserUsecase(userRepository)
    authController := handler.NewAuthHandler(userUsecase)
    authController.Route(r)

    movieRepository := repository.NewMovieRepository(db)
    movieUsecase := usecase.NewMovieUsecase(movieRepository)
    movieController := handler.NewMovieHandler(movieUsecase, userUsecase)
    movieController.Route(r)

    seatRepository := repository.NewSeatRepository(db)
    seatUsecase := usecase.NewSeatUsecase(seatRepository)
    seatController := handler.NewSeatHandler(seatUsecase, userUsecase)
    seatController.Route(r)


    r.Run(":8080")
}