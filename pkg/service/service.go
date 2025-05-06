package service

import (
	"github.com/TolyanchikNeProger/dnnd"
	"github.com/TolyanchikNeProger/dnnd/pkg/repository"
)

type Authorization interface {
	CreateUser(user dnnd.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Menu interface {
}

type MenuItem interface {
}

type Service struct {
	Authorization
	Menu
	MenuItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
