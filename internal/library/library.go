package library

import "errors"

type Book struct {
	ID    int
	Title string
}

type Library struct {
	books          []Book
	availableBooks map[int][]Book
}

func New() *Library {
	return &Library{
		books:          []Book{},
		availableBooks: map[int][]Book{},
	}
}

func (l *Library) FindByID(id int) (*Book, error) {
	for _, b := range l.books {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, errors.New("not found")
}

func (l *Library) AddBook(Book) error {
	return nil
}

func (l *Library) GetBooks(id int) []Book {
	return l.books
}

func (l *Library) Borrow(id int) Book {
	books := l.availableBooks[id]
	return books[0]
}

func (l *Library) Return(book Book) {
	l.availableBooks[book.ID] = append(l.availableBooks[book.ID], book)
}
