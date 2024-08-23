package usecase

import (
	"cinema/internal/model"
	"cinema/internal/repository"
	"cinema/internal/util/security"
	"encoding/json"
	"fmt"
	"os"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

type (
	UserUsecase interface {
		RegisterUser(userToRegister model.RegisterRequest) (err error)
		LoginUser(userToLogin model.LoginRequest) (token string, err error)
		GetUserProfile(token string) (userProfile model.UserProfileResponse, err error)
		CheckRole(token string) (role string, err error)
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

func (uu userUsecase) RegisterUser(userToRegister model.RegisterRequest) (err error) {
	encodedPass, _ := security.HashPassword(userToRegister.Password)
	if err = uu.userRepository.Save(model.Users{
		Username: userToRegister.Username,
		Email:    userToRegister.Email,
		Password: encodedPass,
		RoleID:   userToRegister.RoleID,
	}); err != nil {
		return err
	}

	return nil
}

func (uu userUsecase) LoginUser(userToLogin model.LoginRequest) (token string, err error) {
	user, err := uu.userRepository.FindUserByEmailOrUsername(userToLogin.UsernameOrEmail)

	if err != nil {
		return "", err
	}

	if !security.CheckPasswordHash(userToLogin.Password, user.Password) {
		return "", fmt.Errorf("invalid password")
	}

	generatedToken := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	generatedToken.Claims = jwt_lib.MapClaims{
		"Id":  user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token, err = generatedToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uu userUsecase) GetUserProfile(token string) (model.UserProfileResponse, error) {
    userProfile := model.UserProfileResponse{}

    claims, err := security.ParseToken(token)
    if err != nil {
        return userProfile, fmt.Errorf("invalid token: %w", err)
    }

    var userId uint
    switch id := (*claims)["Id"].(type) {
    case float64:
        userId = uint(id)
    case int:
        userId = uint(id)
    case int64:
        userId = uint(id)
    case json.Number:
        parsedId, err := id.Int64()
        if err != nil {
            return userProfile, fmt.Errorf("invalid token claims: cannot parse user Id, error: %v", err)
        }
        userId = uint(parsedId)
    default:
        return userProfile, fmt.Errorf("invalid token claims: no user Id or unexpected type, claims received: %+v", *claims)
    }

    user, err := uu.userRepository.FindOneUser(userId)
    if err != nil {
        return userProfile, fmt.Errorf("user not found: %w", err)
    }

    userProfile.Username = user.Username
    userProfile.Email = user.Email
	userProfile.Role = user.Role.Name

    return userProfile, nil
}

func (uu userUsecase) CheckRole(token string) (role string, err error) {
	claims, err := security.ParseToken(token)
	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	var userId uint
	switch id := (*claims)["Id"].(type) {
	case float64:
		userId = uint(id)
	case int:
		userId = uint(id)
	case int64:
		userId = uint(id)
	case json.Number:
		parsedId, err := id.Int64()
		if err != nil {
			return "", fmt.Errorf("invalid token claims: cannot parse user Id, error: %v", err)
		}
		userId = uint(parsedId)
	default:
		return "", fmt.Errorf("invalid token claims: no user Id or unexpected type, claims received: %+v", *claims)
	}

	user, err := uu.userRepository.FindOneUser(userId)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}

	return user.Role.Name, nil
}