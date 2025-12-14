// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
)

// define an author type with a name
//

type author struct {
	name string
}

// define a book type with a title and author
//

type book struct {
	id      string
	title   string
	authors []author
}

// define a library type to track a list of books
//
type library struct {
	books        map[string]book
	author_index map[string][]string
}

func (b book) AuthorNames() string {
	names := ""
	if len(b.authors) > 1 {
		for i, a := range b.authors[:len(b.authors)-1] {
			if i > 0 {
				names += ", "
			}
			names += a.name
		}
		names += " and " + b.authors[len(b.authors)-1].name
	} else if len(b.authors) == 1 {
		names = b.authors[0].name
	}
	return names
}

// define addBook to add a book to the library
//

func (l *library) addBook(b book) {
	if l.books == nil {
		l.books = make(map[string]book)
	}
	if b.id == "" {
		b.id = uuid.NewString()
	}
	l.books[b.id] = b
	for _, a := range b.authors {
		if l.author_index == nil {
			l.author_index = make(map[string][]string)
		}
		l.author_index[a.name] = append(l.author_index[a.name], b.id)
	}
}

// define a lookupByAuthor function to find books by an author's name
//

func (l library) lookupByAuthor(name string) []string {
	return l.author_index[name]
}

func main() {
	// create a new library
	//
	var lib library

	// add 2 books to the library by the same author
	//
	lib.addBook(book{
		title: "Book One",
		authors: []author{
			{name: "Author A"},
		},
	})
	lib.addBook(book{
		title: "Book Two",
		authors: []author{
			{name: "Author A"},
			{name: "Author C"},
		},
	})

	// add 1 book to the library by a different author
	//
	lib.addBook(book{
		title: "Book Three",
		authors: []author{
			{name: "Author B"},
		},
	})

	// dump the library with spew
	//
	fmt.Println("Library author index contents:")
	spew.Dump(lib.author_index)

	// Add some more books by Author C
	//
	lib.addBook(book{
		title: "Book Four",
		authors: []author{
			{name: "Author C"},
		},
	})
	lib.addBook(book{
		title: "Book Five",
		authors: []author{
			{name: "Author C"},
			{name: "Author D"},
			{name: "Author Z"},
		},
	})
	lib.addBook(book{
		title: "Book Six",
		authors: []author{
			{name: "Author A"},
			{name: "Author B"},
			{name: "Author C"},
		},
	})

	lib.addBook(book{
		title: "Book Seven",
		authors: []author{
			{name: "Author E"},
			{name: "Author X"},
			{name: "Author C"},
			{name: "Author Z"},
		},
	})

	lib.addBook(book{
		title: "Book Eight",
		authors: []author{
			{name: "Author F"},
		},
	})

	lib.addBook(book{
		title: "Book Nine",
		authors: []author{
			{name: "Author A"},
			{name: "Author F"},
			{name: "Author X"},
		},
	})

	fmt.Println("Updated library author index contents:")
	spew.Dump(lib.author_index)

	// lookup books by known author in the library
	//
	booksByAuthorC := lib.lookupByAuthor("Author C")

	fmt.Println("Books by author results from index")
	spew.Dump(booksByAuthorC)

	// print out the first book's title and its author's name
	//
	for _, book_id := range booksByAuthorC {
		book := lib.books[book_id]
		if len(book.authors) > 1 {
			fmt.Printf("%v (id: %v) is by author(s): %v\n", book.title, book.id, book.AuthorNames())
		} else {
			fmt.Printf("%v (id: %v) is by author: %v\n", book.title, book.id, book.AuthorNames())
		}
	}

}
