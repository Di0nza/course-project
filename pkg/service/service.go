package service

import (
	courseProject "CourseProject"
	"CourseProject/pkg/repository"
)

type Authorization interface {
	CreateUser(user courseProject.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type BrawlersList interface {
	Create(userID int, brawlersList courseProject.BrawlersList) (int, error)
	GetAll(userID int) ([]courseProject.BrawlersListCalc, error)
	Delete(userID int, listsID int) error
}

type Service struct {
	Authorization
	BrawlersList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		BrawlersList:  NewBrawlersListService(repos.BrawlersList),
	}
}
