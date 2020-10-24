package service

import (
	"net/http"

	"api-alkitab/passage"
	"api-alkitab/utils"
)

type PassageService struct {
	passageRepository passage.PassageRepository
}

func (u *PassageService) PassageChapter(passage string, chapter int) (map[string]interface{}, error) {
	response, err := u.passageRepository.PassageChapter(passage, chapter)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, err.Error()), err
	}

	mapResponse := utils.Message(http.StatusOK, "success to fetching data")
	mapResponse["response"] = response

	return mapResponse, nil
}

func (u *PassageService) PassageChapterNumber(passage string, chapter int, number int) (map[string]interface{}, error) {
	response, err := u.passageRepository.PassageChapterNumber(passage, chapter, number)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, err.Error()), err
	}

	mapResponse := utils.Message(http.StatusOK, "success to fetching data")
	mapResponse["response"] = response

	return mapResponse, nil
}

func CreatePassageService(passageRepository passage.PassageRepository) passage.PassageService {
	return &PassageService{passageRepository}
}
