package schema

import "mime/multipart"

type CreateMovies struct {
	Title       string                `validate:"required" form:"title"`
	Description string                `validate:"required" form:"description"`
	Rating      float64               `form:"rating"`
	Image       *multipart.FileHeader `form:"image"`
}

type GetMovies struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type UpdateMovies struct {
	Title       string                `form:"title"`
	Description string                `form:"description"`
	Rating      float64               `form:"rating"`
	Image       *multipart.FileHeader `form:"image"`
}
