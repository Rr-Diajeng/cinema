package handler

import (
	"cinema/internal/model"
	"cinema/internal/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
	userUsecase usecase.UserUsecase
}

func NewAuthHandler (userUsecase usecase.UserUsecase) *AuthHandler{
	return &AuthHandler{
		userUsecase: userUsecase,
	}
}

func (ah *AuthHandler) register(c *gin.Context){
	// bind validate, if error returns 400 bad request
	registerRequest := model.RegisterRequest{}
	err := c.ShouldBindJSON(&registerRequest)

	if err != nil{
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error": strings.Split(err.Error(), "' Error:")[1],
		})
		return
	}

	// call register usecase, if err not nil then return error 500 internal server error
	err = ah.userUsecase.RegisterUser(c, registerRequest)

    if err!= nil{
        c.JSON(500, gin.H{
            "message": "Internal Server Error",
            "error": strings.Split(err.Error(), " violates unique constraint \"uni_users_email\" (SQLSTATE 23505)")[0],
        })
        return
    }

    c.JSON(200, gin.H{
        "message": "User registered successfully",
    })

}

func (ah *AuthHandler) login(c *gin.Context){
	loginRequest := model.LoginRequest{}
	err := c.ShouldBindJSON(loginRequest)

	if err != nil{
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error": strings.Split(err.Error(), "' Error:")[1],
		})
		return
	}

	token, err := ah.userUsecase.LoginUser(c, loginRequest)
	if err != nil{
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully login",
		"token": token,
	})

}

func (ah *AuthHandler) Route(r *gin.Engine) *gin.Engine{
	public := r.Group("/api/auth")

	public.POST("/register", ah.register)
	public.POST("/login", ah.login)

	return r
}