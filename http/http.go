package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/enrichman/coverage/internal/library"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	// middlewares
	r.Use(ErrorHandler)

	library := library.New()
	libraryHandlers(r, library)

	// handler for graceful shutdown
	quit := exitHandler(r)

	srv := &http.Server{
		Addr:    ":8088",
		Handler: r,
	}

	gracefulListenAndServe(srv, quit)
}

func gracefulListenAndServe(srv *http.Server, quit chan os.Signal) {
	go func() {
		err := srv.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("closing server: %s\n", err)
		}
	}()

	// wait for the quit signal
	<-quit
	quitServer(srv)
}

func quitServer(srv *http.Server) {
	log.Println("Shutting down server...")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
