package passage

import "api-alkitab/passage/models"

// PassageRepository is interface to implemented repository
type PassageRepository interface {
	PassageChapter(passage string, chapter int) (*models.Passage, error)
	PassageChapterVerse(passage string, chapter int, verse int) (*models.Passage, error)

	PassageChapterV2(passage string, chapter int, ver string) (*models.PassageV2, error)
	PassageChapterVerseV2(passage string, chapter int, verse int, ver string) (*models.PassageV2, error)
}
