package models

// PassageList is json structure for passage list
type PassageList struct {
	BookNumber   int    `json:"book_number"`
	Abbreviation string `json:"abbreviation"`
	Passage      string `json:"passage"`
	TotalVerse   int    `json:"total_verse"`
}
