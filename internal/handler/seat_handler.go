package handler

import (
	"cinema/internal/model"
	"cinema/internal/usecase"
	"cinema/internal/util/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SeatHandler struct{
	seatUsecase usecase.SeatUsecase
	userUsecase usecase.UserUsecase
}

func NewSeatHandler(seatUsecase usecase.SeatUsecase, userUsecase usecase.UserUsecase) *SeatHandler{
	return &SeatHandler{
		seatUsecase: seatUsecase,
		userUsecase: userUsecase,
	}
}

func (sh *SeatHandler) addSeat(c *gin.Context){
	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		c.JSON(400, gin.H{
			"message": "Authorization token is required",
		})

		return
	}

	if !strings.HasPrefix(authToken, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Authorization token must be in the format Bearer",
		})

		return
	}

	token := strings.Split(authToken, " ")[1]

	if security.BlacklistedTokens[token] {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User needs to log in again",
		})

		return
	}

	role, err := sh.userUsecase.CheckRole(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
			"error":   err,
		})
	}

	if role == "staff" {

		addRequest := model.SeatInput{}
		c.ShouldBindJSON(&addRequest)

		err = sh.seatUsecase.AddSeat(addRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Internal server error",
				"error":   err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"message": "Input seat has been successful",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User doesn't have permission to input movie",
		})
	}
}

func (sh *SeatHandler) Route(r *gin.Engine) *gin.Engine{
	public := r.Group("/api/seat")

	public.POST("/addSeat", sh.addSeat)

	return r
}