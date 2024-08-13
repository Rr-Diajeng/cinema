package usecase

import (
	"cinema/internal/model"
	"cinema/internal/repository"
	"cinema/internal/util/security"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type (
	MovieUsecase interface {
		InputMovie(c context.Context, movieToAdd model.AddMovieRequest) (err error)
		CheckRole(c context.Context, token string) (role string, err error)
	}

	movieUsecase struct {
		movieRepository repository.MovieRepository
		userRepository  repository.UserRepository
	}
)

func NewMovieUsecase(movie repository.MovieRepository, user repository.UserRepository) MovieUsecase {
	return movieUsecase{
		movieRepository: movie,
		userRepository: user,
	}
}

func (mu movieUsecase) InputMovie(c context.Context, movieToAdd model.AddMovieRequest) (err error) {

	var genres []model.Genres

	for _, genreID := range movieToAdd.Genres {
		var genre model.Genres
		if err := mu.movieRepository.FindGenreByID(c, genreID, &genre); err != nil {
			return err
		}

		genres = append(genres, genre)
	}

	duration, err := time.ParseDuration(movieToAdd.Duration)
	if err != nil{
		return err
	}

	if err := mu.movieRepository.AddMovie(c, model.Movies{
		Title:       movieToAdd.Title,
		Genres:      genres,
		Duration:    duration,
		ReleaseDate: movieToAdd.ReleaseDate,
		Synopsis:    movieToAdd.Synopsis,
		BasePrice:   movieToAdd.BasePrice,
	}); err != nil {
		return err
	}

	return nil
}

func (mu movieUsecase) CheckRole(c context.Context, token string) (role string, err error) {
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

	user, err := mu.userRepository.FindOneUser(c, userId)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}

	return user.Role.Name, nil
}
