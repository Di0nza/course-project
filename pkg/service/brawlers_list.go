package service

import (
	courseProject "CourseProject"
	"CourseProject/pkg/repository"
)

type BrawlersListService struct {
	repo repository.BrawlersList
}

func NewBrawlersListService(repo repository.BrawlersList) *BrawlersListService {
	return &BrawlersListService{repo: repo}
}

func (s *BrawlersListService) Create(userId int, brawlersList courseProject.BrawlersList) (int, error) {
	return s.repo.Create(userId, brawlersList)
}

func (s *BrawlersListService) GetAll(userId int) ([]courseProject.BrawlersListCalc, error) {
	return s.repo.GetAll(userId)
}

func (s *BrawlersListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

//func (s *TodoListService) Update(userId, listId int, input todo.UpdateListInput) error {
//	if err := input.Validate(); err != nil {
//		return err
//	}
//
//	return s.repo.Update(userId, listId, input)
//}
