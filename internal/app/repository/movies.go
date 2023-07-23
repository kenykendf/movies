package repository

import (
	"xsis/assignment/internal/app/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Movies struct {
	DB *gorm.DB
}

func NewMoviesRepo(DB *gorm.DB) *Movies {
	return &Movies{
		DB: DB,
	}
}

func (m *Movies) CreateMovies(params model.Movies) error {
	movies := model.Movies{
		Title:       params.Title,
		Rating:      0,
		Description: params.Description,
		Image:       params.Image,
	}

	err := m.DB.Create(&movies).Error
	if err != nil {
		logrus.Error("unable to create movies")
		return err

	}
	return nil
}

func (m *Movies) GetMovies() ([]model.Movies, error) {
	var (
		movies []model.Movies
	)

	err := m.DB.Find(&movies).Error
	if err != nil {
		return movies, err
	}

	return movies, err
}

func (m *Movies) GetMoviesByID(ID int) (model.Movies, error) {
	var (
		movie model.Movies
	)

	err := m.DB.Find(&movie).Where("ID = ?", ID).Error
	if err != nil {
		logrus.Error("unable to retrieve movie detail")
		return movie, err
	}

	return movie, nil
}

func (m *Movies) UpdateMoviesByID(ID int, params model.Movies) error {
	movies := &model.Movies{
		ID:          ID,
		Title:       params.Title,
		Description: params.Description,
		Rating:      params.Rating,
		Image:       params.Image,
	}

	err := m.DB.Updates(&movies).Error
	if err != nil {
		logrus.Error("unable to update product")
		return err
	}

	return nil
}

func (m *Movies) DeleteMoviesByID(ID int) error {
	movie := &model.Movies{
		ID: ID,
	}
	err := m.DB.Delete(&movie).Error

	if err != nil {
		logrus.Error("unable to delete movie")
		return err
	}

	return nil
}
