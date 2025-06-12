package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("rocord not found")
)

type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
