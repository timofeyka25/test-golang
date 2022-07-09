package testGolang

import "time"

type Film struct {
	Id          int       `json:"id" swaggerignore:"true"`
	Title       string    `json:"title" binding:"required"`
	ReleaseDate time.Time `json:"releaseDate" binding:"required" swaggertype:"string" format:"date-time"`
	CreatedAt   time.Time `json:"createdAt" swaggerignore:"true"`
}

type FilmOutput struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"releaseDate"`
	CreatedAt   time.Time `json:"-"`
}
