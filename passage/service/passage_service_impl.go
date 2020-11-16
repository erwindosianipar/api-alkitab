package service

import (
	"net/http"

	"api-alkitab/passage"
	"api-alkitab/passage/models"
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

func (u *PassageService) PassageChapterVerse(passage string, chapter int, verse int) (map[string]interface{}, error) {
	response, err := u.passageRepository.PassageChapterVerse(passage, chapter, verse)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, err.Error()), err
	}

	mapResponse := utils.Message(http.StatusOK, "success to fetching data")
	mapResponse["response"] = response

	return mapResponse, nil
}

func (u *PassageService) PassageChapterV2(passage string, chapter int, ver string) (*models.PassageV2, error) {
	response, err := u.passageRepository.PassageChapterV2(passage, chapter, ver)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *PassageService) PassageChapterVerseV2(passage string, chapter int, verse int, ver string) (*models.PassageV2, error) {
	response, err := u.passageRepository.PassageChapterVerseV2(passage, chapter, verse, ver)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func CreatePassageService(passageRepository passage.PassageRepository) passage.PassageService {
	return &PassageService{passageRepository}
}
