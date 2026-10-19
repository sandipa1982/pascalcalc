package server

import (
	"calculator-app/backend/internal/config"
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
)

// Run starts one independently deployable HTTP service with graceful shutdown.
func Run(name string, defaultPort int, router func(config.Config) *http.ServeMux) error {
	if err := config.LoadDotEnv(); err != nil {
		return err
	}
	cfg, err := config.LoadForService(defaultPort)
	if err != nil {
		return err
	}
	addr := net.JoinHostPort(cfg.ServerHost, strconv.Itoa(cfg.ServerPort))
	srv := &http.Server{
		Addr: addr, Handler: router(cfg), ReadTimeout: cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout, IdleTimeout: cfg.IdleTimeout,
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signals)
	done := make(chan struct{})
	stop := make(chan struct{})
	go func() {
		defer close(done)
		select {
		case <-signals:
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Printf("%s shutdown error: %v", name, err)
			}
		case <-stop:
		}
	}()
	log.Printf("PascalCalc %s listening on %s", name, addr)
	err = srv.ListenAndServe()
	close(stop)
	<-done
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
