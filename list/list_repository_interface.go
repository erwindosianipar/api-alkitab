package list

import "api-alkitab/list/models"

// ListRepository is interface to implemented repository
type ListRepository interface {
	ListPassage() (*[]models.PassageList, error)
}
