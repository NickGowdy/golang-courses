package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) string() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{ID: 1, Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", YearPublished: 1979},
	{ID: 2, Title: "The Hobbit", Author: "J.R.R. Tolkien", YearPublished: 1937},
	{ID: 3, Title: "The 3rd book", Author: "J.R.R. Tolkien", YearPublished: 1937},
	{ID: 4, Title: "The 4th book", Author: "J.R.R. Tolkien", YearPublished: 1937},
	{ID: 5, Title: "The 5th book", Author: "J.R.R. Tolkien", YearPublished: 1937},
	{ID: 6, Title: "Derp 1", Author: "J.R.R. Tolkien", YearPublished: 1937},
	{ID: 7, Title: "Derp 2", Author: "J.R.R. Tolkien", YearPublished: 1937},
	{ID: 8, Title: "Something", Author: "J.R.R. Tolkien", YearPublished: 1937},
	{ID: 9, Title: "Something else", Author: "J.R.R. Tolkien", YearPublished: 1937},
	{ID: 10, Title: "Meh", Author: "J.R.R. Tolkien", YearPublished: 1937},
}
