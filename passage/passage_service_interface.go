package passage

type PassageService interface {
	PassageChapter(passage string, chapter int) (map[string]interface{}, error)
	PassageChapterNumber(passage string, chapter int, number int) (map[string]interface{}, error)
}
