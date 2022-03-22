package models

// Verse json structure for verse
type Verse struct {
	Number int    `xml:"number" json:"number"`
	Text   string `xml:"text" json:"text"`
}

// Verses is json structure for verses
type Verses struct {
	Verse []Verse `xml:"verse" json:"verse"`
}

// Chapter is json structure for chapter
type Chapter struct {
	Chap   int    `xml:"chap" json:"chap"`
	Verses Verses `xml:"verses" json:"verses"`
}

// Book is json structure for chapter
type Book struct {
	BookID  int     `xml:"book_id" json:"book_id"`
	Title   string  `xml:"title" json:"title"`
	Chapter Chapter `xml:"chapter" json:"chapter"`
}

// Passage is json structure for passage
type Passage struct {
	Title string `xml:"title" json:"title"`
	Books []Book `xml:"book" json:"passage"`
}

// VerseV2 json structure for verse v2
type VerseV2 struct {
	Number int    `xml:"number" json:"verse"`
	Text   string `xml:"text" json:"content"`
}

// PassageV2 json structure for passage v2
type PassageV2 struct {
	Title   string    `xml:"title" json:"title"`
	BookID  int       `xml:"book>book_id" json:"book_number"`
	Chapter int       `xml:"book>chapter>chap" json:"chapter"`
	Verses  []VerseV2 `xml:"book>chapter>verses>verse" json:"verses"`
}

type VerseV3 struct {
	Number int    `xml:"number" json:"verse"`
	Title  string `xml:"title" json:"type"`
	Text   string `xml:"text" json:"content"`
}

type PassageV3 struct {
	Title   string    `xml:"title" json:"title"`
	BookID  int       `xml:"book>book_id" json:"book_number"`
	Chapter int       `xml:"book>chapter>chap" json:"chapter"`
	Verses  []VerseV3 `xml:"book>chapter>verses>verse" json:"verses"`
}
