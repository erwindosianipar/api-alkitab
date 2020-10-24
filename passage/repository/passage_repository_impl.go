package repository

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"api-alkitab/passage"
	"api-alkitab/passage/models"

	"github.com/sirupsen/logrus"
)

var _ passage.PassageRepository = &PassageRepositoryImpl{}

type PassageRepositoryImpl struct {
	baseURL string
}

func (u *PassageRepositoryImpl) PassageChapter(passage string, chapter int) (*models.Passage, error) {
	params := fmt.Sprintf("?passage=%s+%v", passage, chapter)

	resp, err := http.Get(u.baseURL + params)
	logrus.Info(resp.Request.URL.String())
	if err != nil {
		logrus.Error("[repository][Passage][Get]", err)
		return nil, errors.New("error to call api")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error("[repository][Passage][ReadAll]", err)
			return nil, errors.New("error read response data")
		}

		data := &models.Passage{}
		err = xml.Unmarshal(bytes, data)
		if err != nil {
			logrus.Error("[repository][Passage][Unmarshal]", err)
			return nil, errors.New("error decode response")
		}

		return data, nil
	}

	logrus.Error("[repository][Passage][statusCode]", resp.StatusCode)
	return nil, errors.New("error something went wrong")
}

func (u *PassageRepositoryImpl) PassageChapterNumber(passage string, chapter int, number int) (*models.Passage, error) {
	params := fmt.Sprintf("?passage=%s+%v:%v", passage, chapter, number)

	resp, err := http.Get(u.baseURL + params)
	logrus.Info(resp.Request.URL.String())
	if err != nil {
		logrus.Error("[repository][PassageChapterNumber][Get]", err)
		return nil, errors.New("error to call api")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error("[repository][PassageChapterNumber][ReadAll]", err)
			return nil, errors.New("error read response data")
		}

		data := &models.Passage{}
		err = xml.Unmarshal(bytes, data)
		if err != nil {
			logrus.Error("[repository][PassageChapterNumber][Unmarshal]", err)
			return nil, errors.New("error decode response")
		}

		return data, nil
	}

	logrus.Error("[repository][PassageChapterNumber][statusCode]", resp.StatusCode)
	return nil, errors.New("error something went wrong")
}

func CreatePassageRepository(baseURL string) passage.PassageRepository {
	return &PassageRepositoryImpl{baseURL}
}
