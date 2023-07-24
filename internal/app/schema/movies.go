package schema

import "mime/multipart"

type QueryParams struct {
	Limit   int
	Offset  int
	Title   string
	SortBy  string
	AscDesc int // value 0 desc
}

type CreateMovies struct {
	Title       string                `validate:"required,min=1,max=100" form:"title"`
	Description string                `validate:"required" form:"description"`
	Rating      float64               `validate:"required,number" form:"rating"`
	Image       *multipart.FileHeader `validate:"image" form:"image"`
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
	Rating      float64               `validate:"number" form:"rating"`
	Image       *multipart.FileHeader `form:"image"`
}
