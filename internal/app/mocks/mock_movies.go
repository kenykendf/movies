// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/service/movies.go

// Package mocks is a generated GoMock package.
package mocks

import (
	multipart "mime/multipart"
	reflect "reflect"
	model "xsis/assignment/internal/app/model"
	schema "xsis/assignment/internal/app/schema"

	gomock "github.com/golang/mock/gomock"
)

// MockMoviesRepo is a mock of MoviesRepo interface.
type MockMoviesRepo struct {
	ctrl     *gomock.Controller
	recorder *MockMoviesRepoMockRecorder
}

// MockMoviesRepoMockRecorder is the mock recorder for MockMoviesRepo.
type MockMoviesRepoMockRecorder struct {
	mock *MockMoviesRepo
}

// NewMockMoviesRepo creates a new mock instance.
func NewMockMoviesRepo(ctrl *gomock.Controller) *MockMoviesRepo {
	mock := &MockMoviesRepo{ctrl: ctrl}
	mock.recorder = &MockMoviesRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMoviesRepo) EXPECT() *MockMoviesRepoMockRecorder {
	return m.recorder
}

// CreateMovies mocks base method.
func (m *MockMoviesRepo) CreateMovies(params model.Movies) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMovies", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMovies indicates an expected call of CreateMovies.
func (mr *MockMoviesRepoMockRecorder) CreateMovies(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMovies", reflect.TypeOf((*MockMoviesRepo)(nil).CreateMovies), params)
}

// DeleteMoviesByID mocks base method.
func (m *MockMoviesRepo) DeleteMoviesByID(ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMoviesByID", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMoviesByID indicates an expected call of DeleteMoviesByID.
func (mr *MockMoviesRepoMockRecorder) DeleteMoviesByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMoviesByID", reflect.TypeOf((*MockMoviesRepo)(nil).DeleteMoviesByID), ID)
}

// GetMovies mocks base method.
func (m *MockMoviesRepo) GetMovies(search schema.QueryParams) ([]model.Movies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies", search)
	ret0, _ := ret[0].([]model.Movies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovies indicates an expected call of GetMovies.
func (mr *MockMoviesRepoMockRecorder) GetMovies(search interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockMoviesRepo)(nil).GetMovies), search)
}

// GetMoviesByID mocks base method.
func (m *MockMoviesRepo) GetMoviesByID(ID int) (model.Movies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMoviesByID", ID)
	ret0, _ := ret[0].(model.Movies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMoviesByID indicates an expected call of GetMoviesByID.
func (mr *MockMoviesRepoMockRecorder) GetMoviesByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMoviesByID", reflect.TypeOf((*MockMoviesRepo)(nil).GetMoviesByID), ID)
}

// UpdateMoviesByID mocks base method.
func (m *MockMoviesRepo) UpdateMoviesByID(ID int, params model.Movies) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMoviesByID", ID, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMoviesByID indicates an expected call of UpdateMoviesByID.
func (mr *MockMoviesRepoMockRecorder) UpdateMoviesByID(ID, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMoviesByID", reflect.TypeOf((*MockMoviesRepo)(nil).UpdateMoviesByID), ID, params)
}

// MockImageUploader is a mock of ImageUploader interface.
type MockImageUploader struct {
	ctrl     *gomock.Controller
	recorder *MockImageUploaderMockRecorder
}

// MockImageUploaderMockRecorder is the mock recorder for MockImageUploader.
type MockImageUploaderMockRecorder struct {
	mock *MockImageUploader
}

// NewMockImageUploader creates a new mock instance.
func NewMockImageUploader(ctrl *gomock.Controller) *MockImageUploader {
	mock := &MockImageUploader{ctrl: ctrl}
	mock.recorder = &MockImageUploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageUploader) EXPECT() *MockImageUploaderMockRecorder {
	return m.recorder
}

// UploadImage mocks base method.
func (m *MockImageUploader) UploadImage(input *multipart.FileHeader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadImage", input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadImage indicates an expected call of UploadImage.
func (mr *MockImageUploaderMockRecorder) UploadImage(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadImage", reflect.TypeOf((*MockImageUploader)(nil).UploadImage), input)
}
