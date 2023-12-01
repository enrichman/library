package library

import (
	"encoding/csv"
	"errors"
	"math/rand"
	"sort"
	"strings"

	_ "embed"
)

//go:embed books.csv
var booksCSV string

type Book struct {
	ISBN   string
	Title  string
	Author string
}

type BookItem struct {
	ID        int
	Book      Book
	Available bool
}

type Library struct {
	catalog        []Book
	availableBooks map[string][]BookItem
}

func New() (*Library, error) {
	reader := csv.NewReader(strings.NewReader(booksCSV))
	reader.Comma = ';'

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	catalog := []Book{}
	availableBooks := map[string][]BookItem{}

	for i, line := range lines {
		if i > 0 {
			book := Book{
				ISBN:   line[0],
				Title:  line[1],
				Author: line[2],
			}

			catalog = append(catalog, book)

			copies := rand.Intn(10)
			for i := 0; i <= copies; i++ {
				items := availableBooks[book.ISBN]
				items = append(items, BookItem{
					ID:   i + 1,
					Book: book,
				})
			}

		}
	}

	sort.Slice(catalog, func(i, j int) bool {
		return catalog[i].ISBN < catalog[j].ISBN
	})

	return &Library{
		catalog:        catalog,
		availableBooks: map[string][]BookItem{},
	}, nil
}

func (l *Library) FindByID(id string) (*Book, error) {
	for _, b := range l.catalog {
		if b.ISBN == id {
			return &b, nil
		}
	}
	return nil, errors.New("not found")
}

func (l *Library) AddBook(Book) error {
	return nil
}

func (l *Library) GetBooks() []Book {
	return l.catalog
}

func (l *Library) Borrow(isbn string) BookItem {
	books := l.availableBooks[isbn]
	return books[0]
}

func (l *Library) Return(item BookItem) {
	l.availableBooks[item.Book.ISBN] = append(l.availableBooks[item.Book.ISBN], item)
}
