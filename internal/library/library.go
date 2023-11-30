package library

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

func (l *Library) FindByID(id int) Book {
	for _, b := range l.books {
		if b.ID == id {
			return b
		}
	}
	return Book{}
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
