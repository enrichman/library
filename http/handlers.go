package http

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/enrichman/coverage/internal/library"
	lib "github.com/enrichman/coverage/internal/library"
	"github.com/gin-gonic/gin"
)

type BookResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func libraryHandlers(r *gin.Engine, library *library.Library) {
	r.GET("/books/:id", func(c *gin.Context) {
		paramID := c.Param("id")

		id, err := strconv.Atoi(paramID)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		book, err := library.FindByID(id)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, BookResponse{
			ID:    book.ID,
			Title: book.Title,
		})
	})

	r.POST("/books", func(c *gin.Context) {
		bookRequest := &BookResponse{}

		err := c.BindJSON(bookRequest)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		book := lib.Book{
			ID:    bookRequest.ID,
			Title: bookRequest.Title,
		}
		err = library.AddBook(book)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, BookResponse{
			ID:    book.ID,
			Title: book.Title,
		})
	})
}

func exitHandler(r *gin.Engine) chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	r.GET("/exit", func(c *gin.Context) {
		quit <- syscall.SIGTERM
	})

	return quit
}
