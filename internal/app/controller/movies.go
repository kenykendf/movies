package controller

import (
	"net/http"
	"strconv"
	"xsis/assignment/internal/app/schema"
	"xsis/assignment/internal/pkg/handler"

	"github.com/gin-gonic/gin"
)

type MoviesService interface {
	CreateMovies(req *schema.CreateMovies) error
	GetMovies() ([]schema.GetMovies, error)
	GetMoviesByID(ID int) (schema.GetMovies, error)
	UpdateMoviesByID(ID int, req *schema.UpdateMovies) error
	DeleteMoviesByID(ID int) error
}

type MoviesController struct {
	service MoviesService
}

func NewMoviesController(service MoviesService) *MoviesController {
	return &MoviesController{service: service}
}

func (mc *MoviesController) Create(ctx *gin.Context) {
	req := &schema.CreateMovies{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := mc.service.CreateMovies(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "unable to create movies")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "new movies created", nil)
}

func (mc *MoviesController) GetMovies(ctx *gin.Context) {
	data, err := mc.service.GetMovies()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "unable to fetch movies from database")
		return
	}

	if data == nil {
		data = make([]schema.GetMovies, 0)
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (mc *MoviesController) GetMoviesByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	movieID, _ := strconv.Atoi(ID)

	data, err := mc.service.GetMoviesByID(movieID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "unable to fetch movies from database")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (mc *MoviesController) UpdateMoviesByID(ctx *gin.Context) {
	var req schema.UpdateMovies
	ID := ctx.Param("id")
	movieID, _ := strconv.Atoi(ID)

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := mc.service.UpdateMoviesByID(movieID, &req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "unable to update movie")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}

func (mc *MoviesController) DeleteMoviesByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	movieID, _ := strconv.Atoi(ID)

	err := mc.service.DeleteMoviesByID(movieID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "unable to delete movie")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}
