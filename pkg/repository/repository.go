package repository

import (
	courseProject "CourseProject"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user courseProject.User) (int, error)
	GetUsername(email, password string) (courseProject.User, error)
}

type BrawlersList interface {
	Create(userId int, brawlersList courseProject.BrawlersList) (int, error)
	GetAll(userId int) ([]courseProject.BrawlersListCalc, error)
	Delete(userID int, listsID int) error
}

type Repository struct {
	Authorization
	BrawlersList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		BrawlersList:  NewBrawlerListPostgres(db),
	}
}
