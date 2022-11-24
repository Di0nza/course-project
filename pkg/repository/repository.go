package repository

import (
	courseProject "CourseProject"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user courseProject.User) (int, error)
}

type BrawlersList interface {
}
type BrawlersItem interface {
}
type Repository interface {
	Authorization
	BrawlersList
	BrawlersItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
