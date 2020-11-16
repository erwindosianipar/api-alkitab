package passage

import "api-alkitab/passage/models"

type PassageService interface {
	PassageChapter(passage string, chapter int) (map[string]interface{}, error)
	PassageChapterVerse(passage string, chapter int, verse int) (map[string]interface{}, error)

	PassageChapterV2(passage string, chapter int, ver string) (*models.PassageV2, error)
	PassageChapterVerseV2(passage string, chapter int, verse int, ver string) (*models.PassageV2, error)
}
