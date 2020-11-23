package models

// PassageList is json structure for passage list
type PassageList struct {
	BookNumber   int    `json:"book_number"`
	Abbreviation string `json:"abbreviation"`
	BookName     string `json:"book_name"`
	TotalChapter int    `json:"total_chapter"`
}
