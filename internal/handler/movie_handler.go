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

	role, err := mh.movieUsecase.CheckRole(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
			"error":   err,
		})
	}

	inputRequest := model.AddMovieRequest{}
	err = c.ShouldBindJSON(&inputRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})

		return
	}

	if role == "staff" {
		err = mh.movieUsecase.InputMovie(inputRequest)
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
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User doesn't have permission to input movie",
		})
	}

}

func (mh *MovieHandler) editMovie(c *gin.Context) {
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

	role, err := mh.movieUsecase.CheckRole(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
			"error":   err,
		})
	}

	if role == "staff" {

		updateRequest := model.UpdateMovieRequest{}
		c.ShouldBindJSON(&updateRequest)

		err = mh.movieUsecase.UpdateMovie(updateRequest)
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
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User doesn't have permission to input movie",
		})
	}

}

func (mh *MovieHandler) getOneMovie(c *gin.Context) {
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

	movieRequest := model.IdMovieRequest{}
	c.ShouldBindJSON(&movieRequest)

	movieDetails, err := mh.movieUsecase.GetOneMovie(movieRequest.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message":      "can get movie details",
		"movie detail": movieDetails,
	})
}

func (mh *MovieHandler) getMovieInSchedule(c *gin.Context) {
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

	exactTimeRequest := model.MovieInScheduleRequest{}
	c.ShouldBindJSON(&exactTimeRequest)

	movies, err := mh.movieUsecase.GetMovieInSchedule(exactTimeRequest.ExactTime)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal Server Error",
			"error": err.Error(),
		})

		return
	}


	c.JSON(200, gin.H{
		"message": "can get the movie",
		"movies": movies,
	})
}

func (mh *MovieHandler) deleteMovie(c *gin.Context){

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

	role, err := mh.movieUsecase.CheckRole(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
			"error":   err,
		})
	}

	if role == "staff" {

		updateRequest := model.IdMovieRequest{}
		c.ShouldBindJSON(&updateRequest)

		err = mh.movieUsecase.DeleteMovieByID(updateRequest.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Internal server error",
				"error":   err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"message": "Movie has been deleted",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User doesn't have permission to input movie",
		})
	}
}

func (mh *MovieHandler) Route(r *gin.Engine) *gin.Engine {
	public := r.Group("/api/movie")

	public.POST("/addMovie", mh.addMovie)
	public.PUT("/editMovie", mh.editMovie)
	public.GET("/getOneMovie", mh.getOneMovie)
	public.GET("/getMovieInSchedule", mh.getMovieInSchedule)
	public.DELETE("/deleteMovie", mh.deleteMovie)

	return r
}
