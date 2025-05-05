package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
