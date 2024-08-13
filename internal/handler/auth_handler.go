package handler

import (
	"cinema/internal/model"
	"cinema/internal/usecase"
	"cinema/internal/util/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userUsecase usecase.UserUsecase
}

func NewAuthHandler(userUsecase usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{
		userUsecase: userUsecase,
	}
}

func (ah *AuthHandler) register(c *gin.Context) {
	registerRequest := model.RegisterRequest{}
	err := c.ShouldBindJSON(&registerRequest)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	err = ah.userUsecase.RegisterUser(c, registerRequest)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User registered successfully",
	})

}

func (ah *AuthHandler) login(c *gin.Context) {
	loginRequest := model.LoginRequest{}
	err := c.ShouldBindJSON(&loginRequest)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	token, err := ah.userUsecase.LoginUser(c, loginRequest)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully login",
		"token":   token,
	})

}

func (ah *AuthHandler) logout(c *gin.Context) {

	authToken := c.GetHeader("Authorization")
	tokenParts := strings.Split(authToken, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Authorization token must be in the format 'Bearer {token}'",
		})

		return
	}

	token := tokenParts[1]

	security.BlacklistedTokens[token] = true

	c.JSON(http.StatusOK, gin.H{
		"message": "User has been logged out",
	})
}

func (ah *AuthHandler) getUserProfile(c *gin.Context) {

	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		c.JSON(400, gin.H{
			"message": "Authorization token is required",
		})

		return
	}

	tokenParts := strings.Split(authToken, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Authorization token must be in the format 'Bearer {token}'",
		})
		return
	}
	token := tokenParts[1]

	if security.BlacklistedTokens[token] {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User needs to log in again",
		})

		return
	}

	userProfile, err := ah.userUsecase.GetUserProfile(c, token)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully get user profile",
		"user":    userProfile,
	})
}

func (ah *AuthHandler) Route(r *gin.Engine) *gin.Engine {
	public := r.Group("/api/auth")

	public.POST("/register", ah.register)
	public.POST("/login", ah.login)
	public.POST("/logout", ah.logout)

	public.GET("/getUserProfile", ah.getUserProfile)

	return r
}
