package repository

import (
	"github.com/TolyanchikNeProger/dnnd"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user dnnd.User) (int, error)
	GetUser(username, password string) (dnnd.User, error)
}

type Menu interface {
}

type MenuItem interface {
}

type Repository struct {
	Authorization
	Menu
	MenuItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
