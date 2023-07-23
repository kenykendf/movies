package service

import (
	"errors"
	"xsis/assignment/internal/app/model"
	"xsis/assignment/internal/app/schema"

	"github.com/sirupsen/logrus"
)

type MoviesRepo interface {
	CreateMovies(params model.Movies) error
	GetMovies() ([]model.Movies, error)
	GetMoviesByID(ID int) (model.Movies, error)
	UpdateMoviesByID(ID int, params model.Movies) error
	DeleteMoviesByID(ID int) error
}

type MoviesService struct {
	moviesRepo MoviesRepo
}

func NewMoviesService(moviesRepo MoviesRepo) *MoviesService {
	return &MoviesService{
		moviesRepo: moviesRepo,
	}
}

func (ms *MoviesService) CreateMovies(req *schema.CreateMovies) error {
	movies := model.Movies{
		Title:       req.Title,
		Description: req.Description,
		Rating:      0,
		// Image: req,
	}

	err := ms.moviesRepo.CreateMovies(movies)
	if err != nil {
		logrus.Error("unable to create new movies")
		return err
	}

	return nil
}

func (ms *MoviesService) GetMovies() ([]schema.GetMovies, error) {
	var response []schema.GetMovies

	data, err := ms.moviesRepo.GetMovies()
	if err != nil {
		return nil, errors.New("movies data not found")
	}

	for _, v := range data {
		var movie schema.GetMovies
		movie.ID = v.ID
		movie.Title = v.Title
		movie.Description = v.Description
		movie.Rating = v.Rating
		movie.Image = v.Image
		movie.CreatedAt = v.CreatedAt.String()
		movie.UpdatedAt = v.UpdatedAt.String()

		response = append(response, movie)
	}

	return response, nil
}

func (ms *MoviesService) GetMoviesByID(ID int) (schema.GetMovies, error) {
	var response schema.GetMovies

	data, err := ms.moviesRepo.GetMoviesByID(ID)
	if err != nil {
		return response, errors.New("movies data not found")
	}

	response.ID = data.ID
	response.Title = data.Title
	response.Description = data.Description
	response.Rating = data.Rating
	response.Image = data.Image
	response.CreatedAt = data.CreatedAt.String()
	response.UpdatedAt = data.UpdatedAt.String()

	return response, nil
}

func (ms *MoviesService) UpdateMoviesByID(ID int, req *schema.UpdateMovies) error {
	var data model.Movies

	movie, err := ms.moviesRepo.GetMoviesByID(ID)
	if err != nil {
		return errors.New("unable to get movie detail")
	}

	data.Title = req.Title
	if req.Title == "" {
		data.Title = movie.Title
	}
	data.Description = req.Description
	if req.Description == "" {
		data.Description = movie.Description
	}
	data.Rating = req.Rating
	// data.Image = req.Image

	err = ms.moviesRepo.UpdateMoviesByID(ID, data)
	if err != nil {
		logrus.Error("update movie by ID : %w", err)
		return errors.New("unable to update movie")
	}

	return nil
}

func (ms *MoviesService) DeleteMoviesByID(ID int) error {
	_, err := ms.moviesRepo.GetMoviesByID(ID)
	if err != nil {
		return errors.New("unable to get movie detail")
	}

	err = ms.moviesRepo.DeleteMoviesByID(ID)
	if err != nil {
		logrus.Error("delete movie by ID : %w", err)
		return errors.New("unable to delete movie")
	}

	return nil

}
