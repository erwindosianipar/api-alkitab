package service

import (
	"api-alkitab/list"
	"api-alkitab/list/models"
)

// ListService to ensure service implemented
type ListService struct {
	listRepository list.ListRepository
}

// ListPassage exported to get data from list passage repository
func (u *ListService) ListPassage() (*[]models.PassageList, error) {
	response, err := u.listRepository.ListPassage()
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateListService exported to initialize service passage service
func CreateListService(listRepository list.ListRepository) list.ListService {
	return &ListService{listRepository}
}
