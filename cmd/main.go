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

    r.Run(":8080")
}