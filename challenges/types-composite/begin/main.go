// challenges/types-composite/begin/main.go
package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
)

// define an author type with a name
//

type Author struct {
	Name string
}

// define a book type with a title and author
//

type Book struct {
	Id      string
	Title   string
	Authors []Author
}

// define a BookLibrary type to track a list of books
//
type BookLibrary struct {
	Books        map[string]Book
	author_index map[string][]string
}

func (b Book) AuthorNames() string {
	names := ""
	if len(b.Authors) > 1 {
		for i, a := range b.Authors[:len(b.Authors)-1] {
			if i > 0 {
				names += ", "
			}
			names += a.Name
		}
		names += " and " + b.Authors[len(b.Authors)-1].Name
	} else if len(b.Authors) == 1 {
		names = b.Authors[0].Name
	}
	return names
	// Another alternative is the following but it does not have the "and" before the last author
	// out := make([]string, len(b.Authors))
	// for i, a := range b.Authors {
	// 	out[i] = a.Name
	// }
	// return strings.Join(out, ", ")
}

// define addBook to add a book to the library
//

func (l *BookLibrary) AddBook(b Book) {
	if l.Books == nil {
		l.Books = make(map[string]Book)
	}
	if b.Id == "" {
		// This should be an error returned in a real system
		log.Println("The ID of a book is mandatory but was not supplied. Will generate a random UUID.")
		b.Id = uuid.NewString()
	}
	l.Books[b.Id] = b
	for _, a := range b.Authors {
		if l.author_index == nil {
			l.author_index = make(map[string][]string)
		}
		l.author_index[a.Name] = append(l.author_index[a.Name], b.Id)
	}
}

// define a lookupByAuthor function to find books by an author's name
//

func (l BookLibrary) LookupByAuthor(name string) []string {
	return l.author_index[name]
}

func main() {
	// create a new library
	//
	var lib BookLibrary

	// add 2 books to the library by the same author
	//
	lib.AddBook(Book{
		Id:    uuid.NewString(),
		Title: "Book One",
		Authors: []Author{
			{Name: "Author A"},
		},
	})
	lib.AddBook(Book{
		Id:    uuid.NewString(),
		Title: "Book Two",
		Authors: []Author{
			{Name: "Author A"},
			{Name: "Author C"},
		},
	})

	// add 1 book to the library by a different author
	//
	lib.AddBook(Book{
		Id:    uuid.NewString(),
		Title: "Book Three",
		Authors: []Author{
			{Name: "Author B"},
		},
	})

	// dump the library with spew
	//
	fmt.Println("Library author index contents:")
	spew.Dump(lib.author_index)

	// Add some more books by Author C
	//
	lib.AddBook(Book{
		Id:    uuid.NewString(),
		Title: "Book Four",
		Authors: []Author{
			{Name: "Author C"},
		},
	})
	lib.AddBook(Book{
		Id:    uuid.NewString(),
		Title: "Book Five",
		Authors: []Author{
			{Name: "Author C"},
			{Name: "Author D"},
			{Name: "Author Z"},
		},
	})
	lib.AddBook(Book{
		Id:    uuid.NewString(),
		Title: "Book Six",
		Authors: []Author{
			{Name: "Author A"},
			{Name: "Author B"},
			{Name: "Author C"},
		},
	})

	lib.AddBook(Book{
		Id:    uuid.NewString(),
		Title: "Book Seven",
		Authors: []Author{
			{Name: "Author E"},
			{Name: "Author X"},
			{Name: "Author C"},
			{Name: "Author Z"},
		},
	})

	lib.AddBook(Book{
		Id:    uuid.NewString(),
		Title: "Book Eight",
		Authors: []Author{
			{Name: "Author F"},
		},
	})

	lib.AddBook(Book{
		Title: "Book Nine",
		Authors: []Author{
			{Name: "Author A"},
			{Name: "Author F"},
			{Name: "Author X"},
		},
	})

	fmt.Println("Updated library author index contents:")
	spew.Dump(lib.author_index)

	// lookup books by known author in the library
	//
	booksByAuthorC := lib.LookupByAuthor("Author C")

	fmt.Println("Books by author results from index")
	spew.Dump(booksByAuthorC)

	// print out the first book's title and its author's name
	//
	for _, book_id := range booksByAuthorC {
		book := lib.Books[book_id]
		if len(book.Authors) > 1 {
			fmt.Printf("%v (id: %v) is by author(s): %v\n", book.Title, book.Id, book.AuthorNames())
		} else {
			fmt.Printf("%v (id: %v) is by author: %v\n", book.Title, book.Id, book.AuthorNames())
		}
	}

}
