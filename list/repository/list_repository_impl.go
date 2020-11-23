package repository

import (
	"errors"
	"strconv"

	"api-alkitab/list"
	"api-alkitab/list/models"
	"api-alkitab/utils"

	"github.com/sirupsen/logrus"
)

var _ list.ListRepository = &ListRepositoryImpl{}

// ListRepositoryImpl to ensure repository implemented
type ListRepositoryImpl struct {
	csvFile string
}

// ListPassage exported to fetch list passage
func (u *ListRepositoryImpl) ListPassage() (*[]models.PassageList, error) {
	records, err := utils.CSVReader(u.csvFile)
	if err != nil {
		return nil, err
	}

	var passage models.PassageList
	var listPassage []models.PassageList

	for _, record := range records {
		bookNumber, err := strconv.Atoi(record[0])
		if err != nil {
			logrus.Info("[repository][list][ListPassage]", err)
			return nil, errors.New(utils.ErrorSomethingWentWrong)
		}

		passage.BookNumber = bookNumber
		passage.Abbreviation = record[1]
		passage.BookName = record[2]

		totalChapter, err := strconv.Atoi(record[3])
		if err != nil {
			logrus.Info("[repository][list][ListPassage]", err)
			return nil, errors.New(utils.ErrorSomethingWentWrong)
		}

		passage.TotalChapter = totalChapter

		listPassage = append(listPassage, passage)
	}

	return &listPassage, nil
}

// CreateListRepository exported to initialize passage repository
func CreateListRepository(csvFile string) list.ListRepository {
	return &ListRepositoryImpl{csvFile}
}
