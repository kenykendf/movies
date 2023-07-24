package service

import (
	"testing"
	"xsis/assignment/internal/app/mocks"
	"xsis/assignment/internal/app/model"
	"xsis/assignment/internal/app/schema"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMoviesService_GetMovies(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	search := schema.QueryParams{
		Limit:   2,
		Offset:  0,
		SortBy:  "id",
		AscDesc: 1,
	}

	mockImageService := mocks.NewMockImageUploader(mockCtrl)
	mockMoviesRepo := mocks.NewMockMoviesRepo(mockCtrl)
	mockMoviesRepo.EXPECT().GetMovies(search).Return([]model.Movies{
		{
			ID:          2,
			Title:       "MR.ROBOT III",
			Description: "netfilx tv series",
			Rating:      3.6,
			Image:       "https://res.cloudinary.com/cloudinarystudydsoken/image/upload/v1690182820/cloudinarystudydsoken/nasgor.jpeg.jpg",
		},
		{
			ID:          4,
			Title:       "MR.ROBOT IV",
			Description: "netfilx tv series",
			Rating:      0,
			Image:       "https://res.cloudinary.com/cloudinarystudydsoken/image/upload/v1690173536/cloudinarystudydsoken/avataaars.png.png",
		},
	}, nil)

	moviesService := NewMoviesService(mockMoviesRepo, mockImageService)
	movies, err := moviesService.GetMovies(schema.QueryParams{
		Limit:   2,
		Offset:  0,
		SortBy:  "id",
		AscDesc: 1,
	})

	total := len(movies)

	assert.Equal(t, total, 2)
	assert.NoError(t, err)
}

func TestMoviesService_GetMoviesByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	movieID := 2

	mockImageService := mocks.NewMockImageUploader(mockCtrl)
	mockMoviesRepo := mocks.NewMockMoviesRepo(mockCtrl)
	mockMoviesRepo.EXPECT().GetMoviesByID(movieID).Return(model.Movies{
		ID:          2,
		Title:       "MR.ROBOT III",
		Description: "netfilx tv series",
		Rating:      3.6,
		Image:       "https://res.cloudinary.com/cloudinarystudydsoken/image/upload/v1690182820/cloudinarystudydsoken/nasgor.jpeg.jpg",
	}, nil)

	moviesService := NewMoviesService(mockMoviesRepo, mockImageService)
	movies, err := moviesService.GetMoviesByID(movieID)

	assert.Equal(t, schema.GetMovies{
		ID:          2,
		Title:       "MR.ROBOT III",
		Description: "netfilx tv series",
		Rating:      3.6,
		Image:       "https://res.cloudinary.com/cloudinarystudydsoken/image/upload/v1690182820/cloudinarystudydsoken/nasgor.jpeg.jpg",
		CreatedAt:   "0001-01-01 00:00:00 +0000 UTC",
		UpdatedAt:   "0001-01-01 00:00:00 +0000 UTC",
	}, movies)
	assert.NoError(t, err)
}

func TestMoviesServiceNotFound_GetMoviesByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	movieID := 1

	mockImageService := mocks.NewMockImageUploader(mockCtrl)
	mockMoviesRepo := mocks.NewMockMoviesRepo(mockCtrl)
	mockMoviesRepo.EXPECT().GetMoviesByID(movieID).Return(model.Movies{
		ID:          1,
		Title:       "",
		Description: "",
		Rating:      0,
		Image:       "",
	}, nil)

	moviesService := NewMoviesService(mockMoviesRepo, mockImageService)
	movies, err := moviesService.GetMoviesByID(movieID)

	assert.Equal(t, schema.GetMovies{
		ID:          1,
		Title:       "",
		Description: "",
		Rating:      0,
		Image:       "",
		CreatedAt:   "0001-01-01 00:00:00 +0000 UTC",
		UpdatedAt:   "0001-01-01 00:00:00 +0000 UTC",
	}, movies)
	assert.NoError(t, err)
}

func TestMoviesService_CreateMovies(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockMoviesRepo(mockCtrl)
	mockImageService := mocks.NewMockImageUploader(mockCtrl)
	moviesService := NewMoviesService(mockRepo, mockImageService)

	req := schema.CreateMovies{
		Title:       "Avengers",
		Description: "The New Avenger",
		Rating:      5.0,
	}

	t.Run("success_create_movies", func(t *testing.T) {
		mockRepo.EXPECT().CreateMovies(req).Return(nil)

		merr := moviesService.CreateMovies(&req)
		assert.NoError(t, merr)
	})
}
