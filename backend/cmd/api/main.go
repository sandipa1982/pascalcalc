package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"calculator-app/backend/internal/config"
	"calculator-app/backend/internal/server"
)

func main() {
	// Load the optional shared environment file.
	if err := config.LoadDotEnv(); err != nil {
		log.Fatalf("environment loading error: %v", err)
	}

	// Read and validate configuration.
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	// Register application routes.
	router := server.NewRouter(cfg)

	address := net.JoinHostPort(
		cfg.ServerHost,
		strconv.Itoa(cfg.ServerPort),
	)

	httpServer := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Graceful shutdown on Ctrl+C or Docker stop.
	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(
		shutdownSignals,
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer signal.Stop(shutdownSignals)

	go func() {
		<-shutdownSignals

		log.Println("Shutting down PascalCalc API...")

		ctx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Printf("shutdown error: %v", err)
		}
	}()

	log.Printf("PascalCalc API listening on %s", address)

	err = httpServer.ListenAndServe()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server error: %v", err)
	}
}
