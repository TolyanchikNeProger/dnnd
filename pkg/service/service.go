package service

import "github.com/TolyanchikNeProger/dnnd/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
