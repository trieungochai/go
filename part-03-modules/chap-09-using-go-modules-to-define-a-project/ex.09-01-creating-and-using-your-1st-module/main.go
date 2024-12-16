package main

import "bookutil/author"

func main() {
	newAuthor := author.NewAuthor("Ngoc-Hai Trieu", "trieungochai.dev@gmail.com")

	chapterTitle := "Introduction to Go Modules"
	chapterContent := "Go modules provide a structured way to manage dependencies and improve code maintainability."
	newAuthor.WriteChapter(chapterTitle, chapterContent)
	newAuthor.ReviewChapter(chapterTitle, chapterContent)
	newAuthor.FinalizeChapter(chapterTitle)
}
