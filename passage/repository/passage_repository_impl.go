package repository

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"api-alkitab/passage"
	"api-alkitab/passage/models"
	"api-alkitab/utils"

	"github.com/sirupsen/logrus"
)

var _ passage.PassageRepository = &PassageRepositoryImpl{}

// PassageRepositoryImpl to ensure repository implemented
type PassageRepositoryImpl struct {
	baseURL string
}

// PassageChapter exported to fetch data passage and chapter
func (u *PassageRepositoryImpl) PassageChapter(passage string, chapter int) (*models.Passage, error) {
	params := fmt.Sprintf("?passage=%s+%v", passage, chapter)

	resp, err := http.Get(u.baseURL + params)
	if err != nil {
		logrus.Error("[repository][Passage][Get]", err)
		return nil, errors.New(utils.ErrorCallAPI)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error("[repository][Passage][ReadAll]", err)
			return nil, errors.New(utils.ErrorReadResponseData)
		}

		data := &models.Passage{}
		err = xml.Unmarshal(bytes, data)
		if err != nil {
			logrus.Error("[repository][Passage][Unmarshal]", err)
			return nil, errors.New(utils.ErrorDecodeResponse)
		}

		return data, nil
	}

	logrus.Error("[repository][Passage][statusCode]", resp.StatusCode)
	return nil, errors.New(utils.ErrorSomethingWentWrong)
}

// PassageChapterVerse exported to fetch data passage chapter and verse
func (u *PassageRepositoryImpl) PassageChapterVerse(passage string, chapter int, verse int) (*models.Passage, error) {
	params := fmt.Sprintf("?passage=%s+%v:%v", passage, chapter, verse)

	resp, err := http.Get(u.baseURL + params)
	if err != nil {
		logrus.Error("[repository][PassageChapterVerse][Get]", err)
		return nil, errors.New(utils.ErrorCallAPI)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error("[repository][PassageChapterVerse][ReadAll]", err)
			return nil, errors.New(utils.ErrorReadResponseData)
		}

		data := &models.Passage{}
		err = xml.Unmarshal(bytes, data)
		if err != nil {
			logrus.Error("[repository][PassageChapterVerse][Unmarshal]", err)
			return nil, errors.New(utils.ErrorDecodeResponse)
		}

		return data, nil
	}

	logrus.Error("[repository][PassageChapterVerse][statusCode]", resp.StatusCode)
	return nil, errors.New(utils.ErrorSomethingWentWrong)
}

// PassageChapterV2 exported to fetch data passage and chapter for v2
func (u *PassageRepositoryImpl) PassageChapterV2(passage string, chapter int, ver string) (*models.PassageV2, error) {
	params := fmt.Sprintf("?passage=%s+%v", passage, chapter)
	if ver != "" {
		params = fmt.Sprintf("?passage=%s+%v&ver=%s", passage, chapter, ver)
	}

	resp, err := http.Get(u.baseURL + params)
	if err != nil {
		logrus.Error("[repository][PassageChapterV2][Get]", err)
		return nil, errors.New(utils.ErrorCallAPI)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error("[repository][PassageChapterV2][ReadAll]", err)
			return nil, errors.New(utils.ErrorReadResponseData)
		}

		data := &models.PassageV2{}
		err = xml.Unmarshal(bytes, data)
		if err != nil {
			logrus.Error("[repository][PassageChapterV2][Unmarshal]", err)
			return nil, errors.New(utils.ErrorDecodeResponse)
		}

		return data, nil
	}

	logrus.Error("[repository][PassageChapterV2][statusCode]", resp.StatusCode)
	return nil, errors.New(utils.ErrorSomethingWentWrong)
}

// PassageChapterVerseV2 exported to fetch data passage chapter and verse for v2
func (u *PassageRepositoryImpl) PassageChapterVerseV2(passage string, chapter int, verse int, ver string) (*models.PassageV2, error) {
	params := fmt.Sprintf("?passage=%s+%v:%v", passage, chapter, verse)
	if ver != "" {
		params = fmt.Sprintf("?passage=%s+%v:%v&ver=%s", passage, chapter, verse, ver)
	}

	resp, err := http.Get(u.baseURL + params)
	if err != nil {
		logrus.Error("[repository][PassageChapterVerseV2][Get]", err)
		return nil, errors.New(utils.ErrorCallAPI)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error("[repository][PassageChapterVerseV2][ReadAll]", err)
			return nil, errors.New(utils.ErrorReadResponseData)
		}

		data := &models.PassageV2{}
		err = xml.Unmarshal(bytes, data)
		if err != nil {
			logrus.Error("[repository][PassageChapterVerseV2][Unmarshal]", err)
			return nil, errors.New(utils.ErrorDecodeResponse)
		}

		return data, nil
	}

	logrus.Error("[repository][PassageChapterVerseV2][statusCode]", resp.StatusCode)
	return nil, errors.New(utils.ErrorSomethingWentWrong)
}

// PassageChapterV3 exported to fetch data passage and chapter for v3
func (u *PassageRepositoryImpl) PassageChapterV3(passage string, chapter int, ver string) (*models.PassageV3, error) {
	params := fmt.Sprintf("?passage=%s+%v", passage, chapter)
	if ver != "" {
		params = fmt.Sprintf("?passage=%s+%v&ver=%s", passage, chapter, ver)
	}

	resp, err := http.Get(u.baseURL + params)
	if err != nil {
		logrus.Error("[repository][PassageChapterV3][Get]", err)
		return nil, errors.New(utils.ErrorCallAPI)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error("[repository][PassageChapterV3][ReadAll]", err)
			return nil, errors.New(utils.ErrorReadResponseData)
		}

		data := &models.PassageV3{}
		err = xml.Unmarshal(bytes, data)
		if err != nil {
			logrus.Error("[repository][PassageChapterV3][Unmarshal]", err)
			return nil, errors.New(utils.ErrorDecodeResponse)
		}

		var response []models.VerseV3
		for _, v := range data.Verses {
			if len(v.Title) > 0 {
				response = append(response, models.VerseV3{
					Number: 0,
					Title:  "title",
					Text:   v.Title,
				})
			}
			response = append(response, models.VerseV3{
				Number: v.Number,
				Title:  "text",
				Text:   v.Text,
			})
		}

		resp := models.PassageV3{
			Title:   data.Title,
			BookID:  data.BookID,
			Chapter: data.Chapter,
			Verses:  response,
		}

		return &resp, nil
	}

	logrus.Error("[repository][PassageChapterV3][statusCode]", resp.StatusCode)
	return nil, errors.New(utils.ErrorSomethingWentWrong)
}

// CreatePassageRepository exported to initialize passage repository
func CreatePassageRepository(baseURL string) passage.PassageRepository {
	return &PassageRepositoryImpl{baseURL}
}
