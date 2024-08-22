package usecase

import (
	"cinema/internal/model"
	"cinema/internal/repository"
	"cinema/internal/util/security"
	"encoding/json"
	"fmt"
	"time"
)

type (
	MovieUsecase interface {
		InputMovie(movieToAdd model.AddMovieRequest) (err error)
		CheckRole(token string) (role string, err error)
		UpdateMovie(movieToUpdate model.UpdateMovieRequest) error
		GetOneMovie(id uint) (model.OneMovieResponse, error)
	}

	movieUsecase struct {
		movieRepository repository.MovieRepository
		userRepository  repository.UserRepository
	}
)

func NewMovieUsecase(movie repository.MovieRepository, user repository.UserRepository) MovieUsecase {
	return movieUsecase{
		movieRepository: movie,
		userRepository:  user,
	}
}

func (mu movieUsecase) InputMovie(movieToAdd model.AddMovieRequest) (err error) {

	var genres []model.Genres

	for _, genreID := range movieToAdd.Genres {
		var genre model.Genres
		if err := mu.movieRepository.FindGenreByID(genreID, &genre); err != nil {
			return err
		}

		genres = append(genres, genre)
	}

	duration, err := time.ParseDuration(movieToAdd.Duration)
	if err != nil {
		return err
	}

	durationInSeconds := int64(duration.Seconds())

	if err := mu.movieRepository.AddMovie(model.Movies{
		Title:       movieToAdd.Title,
		Genres:      genres,
		Duration:    durationInSeconds,
		ReleaseDate: movieToAdd.ReleaseDate,
		Synopsis:    movieToAdd.Synopsis,
		BasePrice:   movieToAdd.BasePrice,
	}); err != nil {
		return err
	}

	return nil
}

func (mu movieUsecase) CheckRole(token string) (role string, err error) {
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

	user, err := mu.userRepository.FindOneUser(userId)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}

	return user.Role.Name, nil
}

func (mu movieUsecase) UpdateMovie(movieToUpdate model.UpdateMovieRequest) error{
	var genres []model.Genres
	for _, genreID := range movieToUpdate.Genres{
		var genre model.Genres
		if err := mu.movieRepository.FindGenreByID(genreID, &genre); err != nil{
			return err
		}

		genres = append(genres, genre)
	}

	duration, err := time.ParseDuration(movieToUpdate.Duration)
	if err != nil{
		return err
	}

	durationInSeconds := int64(duration.Seconds())

	existingMovie, err := mu.movieRepository.FindMovieByID(movieToUpdate.ID)
	if err != nil{
		return err
	}

	if err := mu.movieRepository.ClearGenres(existingMovie.ID); err != nil{
		return err
	}

	existingMovie.Title = movieToUpdate.Title
    existingMovie.Genres = genres
    existingMovie.Duration = durationInSeconds
    existingMovie.ReleaseDate = movieToUpdate.ReleaseDate
    existingMovie.Synopsis = movieToUpdate.Synopsis
    existingMovie.BasePrice = movieToUpdate.BasePrice

	if err := mu.movieRepository.UpdateMovie(existingMovie); err != nil {
        return err
    }

    return nil
}

func (mu movieUsecase) GetOneMovie(id uint) (model.OneMovieResponse, error){

	movieDetails := model.OneMovieResponse{}

	movie, err :=  mu.movieRepository.FindMovieByID(id)

	if err != nil{
		return movieDetails, fmt.Errorf("movie not found: %w", err)
	}

	var genres []model.Genres
    for _, genre := range movie.Genres {
        var fetchedGenre model.Genres
        if err := mu.movieRepository.FindGenreByID(genre.ID, &fetchedGenre); err != nil {
            return movieDetails, err
        }
        genres = append(genres, fetchedGenre)
    }

    movieDetails.Title = movie.Title
    movieDetails.Genres = genres 
    movieDetails.Duration = movie.Duration
    movieDetails.ReleaseDate = movie.ReleaseDate
    movieDetails.Synopsis = movie.Synopsis
    movieDetails.BasePrice = movie.BasePrice

	return movieDetails, nil
}