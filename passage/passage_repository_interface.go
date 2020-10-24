package passage

import "api-alkitab/passage/models"

type PassageRepository interface {
	PassageChapter(passage string, chapter int) (*models.Passage, error)
	PassageChapterNumber(passage string, chapter int, number int) (*models.Passage, error)
}
