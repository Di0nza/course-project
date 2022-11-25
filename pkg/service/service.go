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
}
type BrawlersItem interface {
}
type Service interface {
	Authorization
	BrawlersList
	BrawlersItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
