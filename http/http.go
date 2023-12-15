package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/enrichman/library/internal/library"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	// middlewares
	r.Use(ErrorHandler)

	library, err := library.New()
	if err != nil {
		log.Panic(err)
	}

	libraryHandlers(r, library)

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8088"
	}

	gracefulListenAndServe(r, port)
}

func gracefulListenAndServe(r *gin.Engine, port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// handler for graceful shutdown
	quit := exitHandler(r)

	go func() {
		log.Printf("server listening at http://localhost%s\n", srv.Addr)

		err := srv.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("closing server: %s\n", err)
		}
	}()

	// wait for the quit signal
	<-quit

	quitServer(srv)
}

func exitHandler(r *gin.Engine) chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	r.GET("/exit", func(c *gin.Context) {
		quit <- syscall.SIGTERM
	})

	return quit
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
