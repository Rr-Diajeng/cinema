package usecase

import (
	"cinema/internal/model"
	"cinema/internal/repository"
	"cinema/internal/util/security"
	"context"
	"os"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

type (
	UserUsecase interface {
		RegisterUser(c context.Context, userToRegister model.RegisterRequest) (err error)
		LoginUser(c context.Context, userToLogin model.LoginRequest) (token string, err error)
		GetUserProfile(c context.Context, idFromToken string) (userProfile model.UserProfileResponse, err error)
	}

	userUsecase struct {
		userRepository repository.UserRepository
	}
)

func NewUserUsecase(user repository.UserRepository) UserUsecase {
	return userUsecase{
		userRepository: user,
	}
}

func (uu userUsecase) RegisterUser(c context.Context, userToRegister model.RegisterRequest) (err error) {
	encodedPass, _ := security.HashPassword(userToRegister.Password)
	if err = uu.userRepository.Save(c, model.Users{
		Username: userToRegister.Username,
		Email:    userToRegister.Email,
		Password: encodedPass,
		RoleID:   userToRegister.RoleID,
	}); err != nil {
		return err
	}

	return nil
}

func (uu userUsecase) LoginUser(c context.Context, userToLogin model.LoginRequest) (token string, err error) {
	user, err := uu.userRepository.FindUserByEmailOrUsername(c, userToLogin.UsernameOrEmail)

	if err != nil {
		return "", err
	}

	if security.CheckPasswordHash(userToLogin.Password, user.Password) {
		generatedToken := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
		generatedToken.Claims = jwt_lib.MapClaims{
			"Id": user.ID,
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		}

		token, err = generatedToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

		if err != nil{
			return "", err
		}
	}

	return token, nil
}

func (uu userUsecase) GetUserProfile(c context.Context, idFromToken string) (userProfile model.UserProfileResponse, err error) {
	return 
}
