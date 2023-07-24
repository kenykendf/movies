package controller

import (
	"net/http"
	"strconv"
	"xsis/assignment/internal/app/schema"
	"xsis/assignment/internal/pkg/handler"
	"xsis/assignment/internal/pkg/reason"

	"github.com/gin-gonic/gin"
)

type MoviesService interface {
	CreateMovies(req *schema.CreateMovies) error
	GetMovies(search schema.QueryParams) ([]schema.GetMovies, error)
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
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.MoviesCreateErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "new movies created", nil)
}

func (mc *MoviesController) GetMovies(ctx *gin.Context) {
	search := schema.QueryParams{}
	search.Limit = ctx.GetInt("limit")
	search.Offset = ctx.GetInt("offset")
	search.Title = ctx.Query("title")
	search.SortBy = ctx.Query("sort_by")
	search.AscDesc = ctx.GetInt("asc")

	data, err := mc.service.GetMovies(search)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.MoviesListsErr)
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
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.MoviesListByIDErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (mc *MoviesController) UpdateMoviesByID(ctx *gin.Context) {
	req := &schema.UpdateMovies{}
	ID := ctx.Param("id")
	movieID, _ := strconv.Atoi(ID)

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := mc.service.UpdateMoviesByID(movieID, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.MoviesUpdateErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}

func (mc *MoviesController) DeleteMoviesByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	movieID, _ := strconv.Atoi(ID)

	err := mc.service.DeleteMoviesByID(movieID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.MoviesDeleteErr)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}
