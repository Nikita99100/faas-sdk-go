package framework

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	writeTimeout    = time.Second * 15
	readTimeout     = time.Second * 15
	gracefulTimeout = time.Second * 15
)

func Start(port int, function http.HandlerFunc) error {
	router := http.NewServeMux()
	router.HandleFunc("/", function)
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		Handler:        router,
	}

	go func() {
		err := srv.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Print("HTTP server closed")
		} else {
			log.Fatalf("can't start listen: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	sig := <-c

	ctx, cancel := context.WithTimeout(context.Background(), gracefulTimeout)
	defer cancel()

	log.Printf("Received signal %s - shutting down...", sig.String())

	_ = srv.Shutdown(ctx)

	log.Print("shutting down")

	return nil
}
