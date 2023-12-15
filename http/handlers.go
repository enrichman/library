package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enrichman/library/internal/library"
	"github.com/gin-gonic/gin"
)

type BookResponse struct {
	ISBN  string `json:"isbn"`
	Title string `json:"title"`
}

type BookListResponse []BookResponse

func libraryHandlers(r *gin.Engine, library *library.Library) {
	r.GET("/books/:id", getBookByID(library))
	r.GET("/books", getBooks(library))
	r.GET("/books/:id/borrow", borrowItem(library))
	r.GET("/books/:id/return/:item_id", returnItem(library))
}

func getBookByID(library *library.Library) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		book, err := library.FindByID(id)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, BookResponse{
			ISBN:  book.ISBN,
			Title: book.Title,
		})
	}
}

func getBooks(library *library.Library) gin.HandlerFunc {
	return func(c *gin.Context) {
		catalog := library.GetBooks()

		limit := 10
		limitStr := c.Query("limit")
		if limitStr != "" {
			limitVal, err := strconv.Atoi(limitStr)
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			if limitVal > 0 && limitVal < 20 {
				limit = limitVal
			}
		}

		catalog = catalog[0:limit]

		c.JSON(http.StatusOK, catalog)
	}
}

func borrowItem(library *library.Library) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		book, err := library.FindByID(id)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		item := library.Borrow(book.ISBN)
		c.JSON(http.StatusOK, item)
	}
}

func returnItem(lib *library.Library) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		itemID := c.Param("item_id")
		fmt.Println(itemID)

		book, err := lib.FindByID(id)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		bookItem := library.BookItem{ID: 383, Book: *book}
		lib.Return(bookItem)
		c.JSON(http.StatusOK, bookItem)
	}
}
