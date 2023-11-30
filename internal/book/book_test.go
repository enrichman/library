package book_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/enrichman/coverage/book"
)

func TestBook(t *testing.T) {
	age, _ := book.ParseAge(23)
	user := book.User{Age: age}

	err := book.Book(user)
	assert.NoError(t, err)
}
