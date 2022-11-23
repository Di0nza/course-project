package service

import "CourseProject/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
