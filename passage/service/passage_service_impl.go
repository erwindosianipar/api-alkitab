package service

import (
	"net/http"

	"api-alkitab/passage"
	"api-alkitab/passage/models"
	"api-alkitab/utils"
)

// PassageService to ensure service implemented
type PassageService struct {
	passageRepository passage.PassageRepository
}

// PassageChapter exported to get data from passage and chapter repository
func (u *PassageService) PassageChapter(passage string, chapter int) (map[string]interface{}, error) {
	response, err := u.passageRepository.PassageChapter(passage, chapter)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, err.Error()), err
	}

	mapResponse := utils.Message(http.StatusOK, "success to fetching data")
	mapResponse["response"] = response

	return mapResponse, nil
}

// PassageChapterVerse exported to get data from passage chapter and verse repository
func (u *PassageService) PassageChapterVerse(passage string, chapter int, verse int) (map[string]interface{}, error) {
	response, err := u.passageRepository.PassageChapterVerse(passage, chapter, verse)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, err.Error()), err
	}

	mapResponse := utils.Message(http.StatusOK, "success to fetching data")
	mapResponse["response"] = response

	return mapResponse, nil
}

// PassageChapterV2 exported to get data from passage and chapter v2 repository
func (u *PassageService) PassageChapterV2(passage string, chapter int, ver string) (*models.PassageV2, error) {
	response, err := u.passageRepository.PassageChapterV2(passage, chapter, ver)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PassageChapterVerseV2 exported to get data from passage chapter and verse v2 repository
func (u *PassageService) PassageChapterVerseV2(passage string, chapter int, verse int, ver string) (*models.PassageV2, error) {
	response, err := u.passageRepository.PassageChapterVerseV2(passage, chapter, verse, ver)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreatePassageService exported to initialize service passage service
func CreatePassageService(passageRepository passage.PassageRepository) passage.PassageService {
	return &PassageService{passageRepository}
}
