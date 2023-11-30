package http

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/enrichman/coverage/internal/library"
	"github.com/gin-gonic/gin"
)

type BookResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func libraryHandlers(r *gin.Engine, library *library.Library) {
	r.GET("/book/:id", func(c *gin.Context) {
		paramID := c.Param("id")

		id, err := strconv.Atoi(paramID)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		book := library.FindByID(id)
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
