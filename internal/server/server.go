package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/smarulanda97/assesment-golang/internal/database"
	"github.com/smarulanda97/assesment-golang/internal/repository"
)

type Config struct {
	Port      string
	JWTSecret string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *chi.Mux
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}

	broker := &Broker{
		config: config,
		router: chi.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(mux *chi.Mux) http.Handler) error {
	server := &http.Server{
		Handler:           binder(b.router),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Addr:              b.config.Port,
	}

	repo, err := database.NewMemoryRepository()
	if err != nil {
		return err
	}

	repository.SetRepository(repo)

	fmt.Printf("starting server on port %s", b.config.Port)

	return server.ListenAndServe()
}
