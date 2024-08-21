package handler

import (
	"cinema/internal/model"
	"cinema/internal/usecase"
	"cinema/internal/util/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieUsecase usecase.MovieUsecase
}

func NewMovieHandler(movieUsecase usecase.MovieUsecase) *MovieHandler {
	return &MovieHandler{
		movieUsecase: movieUsecase,
	}
}

func (mh *MovieHandler) addMovie(c *gin.Context) {

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

	role, err := mh.movieUsecase.CheckRole(c, token)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
			"error": err,
		})
	}

	inputRequest := model.AddMovieRequest{}
	err = c.ShouldBindJSON(&inputRequest)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error": err.Error(),
		})

		return
	}

	if role == "staff"{
		err = mh.movieUsecase.InputMovie(c, inputRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Internal server error",
				"error":   err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"message": "Input movie has been successful",
		})
	} else{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User doesn't have permission to input movie",
		})
	}

}

func (mh *MovieHandler) editMovie(c *gin.Context){
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

	role, err := mh.movieUsecase.CheckRole(c, token)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
			"error": err,
		})
	}

	if role == "staff"{

		updateRequest := model.UpdateMovieRequest{}
		c.ShouldBindJSON(&updateRequest)

		err = mh.movieUsecase.UpdateMovie(c, updateRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Internal server error",
				"error":   err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"message": "Input movie has been successful",
		})
	} else{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User doesn't have permission to input movie",
		})
	}


}

func (mh *MovieHandler) Route(r *gin.Engine) *gin.Engine {
	public := r.Group("/api/movie")

	public.POST("/addMovie", mh.addMovie)
	public.PUT("/editMovie", mh.editMovie)

	return r
}
