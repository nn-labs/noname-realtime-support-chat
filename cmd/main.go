package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"log"
	"net/http"
	"noname-realtime-support-chat/config"
	"noname-realtime-support-chat/internal/health"
	"noname-realtime-support-chat/pkg/logger"
)

func main() {
	cfg, err := config.Get(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Init logger
	newLogger, err := logger.NewLogger(cfg.Environment)
	if err != nil {
		log.Fatalf("can't create logger: %v", err)
	}

	zapLogger, err := newLogger.SetupZapLogger()
	if err != nil {
		log.Fatalf("can't setup zap logger: %v", err)
	}
	defer func(zapLogger *zap.SugaredLogger) {
		err := zapLogger.Sync()
		if err != nil {
			log.Fatalf("can't setup zap logger: %v", err)
		}
	}(zapLogger)

	// Set-up Route
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Handlers
	healthHandler := health.NewHandler()

	router.Route("/api/v1", func(r chi.Router) {
		healthHandler.SetupRoutes(r)
	})

	// Start App
	zapLogger.Infof("Starting HTTP server on port: %v", 5000)
	err = http.ListenAndServe(cfg.PORT, router)
	if err != nil {
		fmt.Println(err)
		zapLogger.Fatalf("Failed to start HTTP server: %v", err)
		return
	}
}
