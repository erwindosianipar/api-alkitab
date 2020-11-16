package models

type Verse struct {
	Number int    `xml:"number" json:"number"`
	Text   string `xml:"text" json:"text"`
}

type Verses struct {
	Verse []Verse `xml:"verse" json:"verse"`
}

type Chapter struct {
	Chap   int    `xml:"chap" json:"chap"`
	Verses Verses `xml:"verses" json:"verses"`
}

type Book struct {
	BookID  int     `xml:"book_id" json:"book_id"`
	Title   string  `xml:"title" json:"title"`
	Chapter Chapter `xml:"chapter" json:"chapter"`
}

type Passage struct {
	Title string `xml:"title" json:"title"`
	Books []Book `xml:"book" json:"passage"`
}

type VerseV2 struct {
	Number int    `xml:"number" json:"verse"`
	Text   string `xml:"text" json:"content"`
}

type PassageV2 struct {
	Title   string    `xml:"title" json:"title"`
	BookID  int       `xml:"book>book_id" json:"book_number"`
	Chapter int       `xml:"book>chapter>chap" json:"chapter"`
	Verses  []VerseV2 `xml:"book>chapter>verses>verse" json:"verses"`
}
