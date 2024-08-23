package usecase

import (
	"cinema/internal/model"
	"cinema/internal/repository"
	"fmt"
	"time"
)

type (
	MovieUsecase interface {
		InputMovie(movieToAdd model.AddMovieRequest) (err error)
		UpdateMovie(movieToUpdate model.UpdateMovieRequest) error
		GetOneMovie(id uint) (model.MovieResponse, error)
		GetMovieInSchedule(exactTime time.Time) ([]model.MovieResponse, error)
		DeleteMovieByID(id uint) error
	}

	movieUsecase struct {
		movieRepository repository.MovieRepository
	}
)

func NewMovieUsecase(movie repository.MovieRepository) MovieUsecase {
	return movieUsecase{
		movieRepository: movie,
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
		StartDate:   movieToAdd.StartDate,
		EndDate:     movieToAdd.EndDate,
	}); err != nil {
		return err
	}

	return nil
}

func (mu movieUsecase) UpdateMovie(movieToUpdate model.UpdateMovieRequest) error {
	var genres []model.Genres
	for _, genreID := range movieToUpdate.Genres {
		var genre model.Genres
		if err := mu.movieRepository.FindGenreByID(genreID, &genre); err != nil {
			return err
		}

		genres = append(genres, genre)
	}

	duration, err := time.ParseDuration(movieToUpdate.Duration)
	if err != nil {
		return err
	}

	durationInSeconds := int64(duration.Seconds())

	existingMovie, err := mu.movieRepository.FindMovieByID(movieToUpdate.ID)
	if err != nil {
		return err
	}

	if err := mu.movieRepository.ClearGenres(existingMovie.ID); err != nil {
		return err
	}

	existingMovie.Title = movieToUpdate.Title
	existingMovie.Genres = genres
	existingMovie.Duration = durationInSeconds
	existingMovie.ReleaseDate = movieToUpdate.ReleaseDate
	existingMovie.Synopsis = movieToUpdate.Synopsis
	existingMovie.BasePrice = movieToUpdate.BasePrice
	existingMovie.StartDate = movieToUpdate.StartDate
	existingMovie.EndDate = movieToUpdate.EndDate

	if err := mu.movieRepository.UpdateMovie(existingMovie); err != nil {
		return err
	}

	return nil
}

func (mu movieUsecase) GetOneMovie(id uint) (model.MovieResponse, error) {

	movieDetails := model.MovieResponse{}

	movie, err := mu.movieRepository.FindMovieByID(id)

	if err != nil {
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
	movieDetails.StartDate = movie.StartDate
	movieDetails.EndDate = movie.EndDate

	return movieDetails, nil
}

func (mu movieUsecase) GetMovieInSchedule(exactTime time.Time) ([]model.MovieResponse, error) {

	movies, err := mu.movieRepository.GetMovieInSchedule(exactTime)

	if err != nil {
		return nil, err
	}

	var movieResponses []model.MovieResponse
	for _, movie := range movies {
		movieResponses = append(movieResponses, model.MovieResponse{
			Title:       movie.Title,
			Genres:      movie.Genres,
			Duration:    movie.Duration,
			ReleaseDate: movie.ReleaseDate,
			Synopsis:    movie.Synopsis,
			BasePrice:   movie.BasePrice,
			StartDate:   movie.StartDate,
			EndDate:     movie.EndDate,
		})
	}

	return movieResponses, nil
}

func (mu movieUsecase) DeleteMovieByID(id uint) error {
	err := mu.movieRepository.DeleteMovieByID(id)

	if err != nil {
		return err
	}

	return nil
}
