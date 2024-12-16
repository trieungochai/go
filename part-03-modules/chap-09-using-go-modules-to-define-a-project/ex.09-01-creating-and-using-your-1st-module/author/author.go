package author

import "fmt"

// represents an author of a book
type Author struct {
	Name    string
	Contact string
}

// creates a new Author instance
func NewAuthor(name, contact string) *Author {
	return &Author{Name: name, Contact: contact}
}

// writes a book chapter by author
func (a *Author) WriteChapter(chapterTitle, content string) {
	fmt.Printf("Author %s is writing a chapter titled '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

// allows the author to review and edit a chapter
func (a *Author) ReviewChapter(chapterTitle, content string) {
	fmt.Printf("Author %s is reviewing a chapter titled '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

// marks the chapter as finalized by the author
func (a *Author) FinalizeChapter(chapterTitle string) {
	fmt.Printf("Author %s has finalized the chapter titled '%s'\n", a.Name, chapterTitle)
}
