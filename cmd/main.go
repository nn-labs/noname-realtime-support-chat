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
	"noname-realtime-support-chat/internal/support"
	"noname-realtime-support-chat/pkg/logger"
	"noname-realtime-support-chat/pkg/mongodb"
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

	// Connect to database
	db, ctx, cancel, err := mongodb.NewConnection(cfg)
	if err != nil {
		zapLogger.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer mongodb.Close(db, ctx, cancel)

	// Ping db
	err = mongodb.Ping(db, ctx)
	if err != nil {
		log.Fatal(err)
	}
	zapLogger.Info("DB connected successfully")

	// Repositories
	supportRepository, err := support.NewRepository(db, zapLogger)
	if err != nil {
		zapLogger.Fatalf("failde to create support repository: %v", err)
	}

	// Services
	supportService, err := support.NewService(supportRepository, zapLogger, &cfg.Salt)
	if err != nil {
		zapLogger.Fatalf("failde to create support service: %v", err)
	}

	// Handlers
	healthHandler := health.NewHandler()
	supportHandler, err := support.NewHandler(supportService)
	if err != nil {
		zapLogger.Fatalf("failde to create support handler: %v", err)
	}

	router.Route("/api/v1", func(r chi.Router) {
		healthHandler.SetupRoutes(r)
		supportHandler.SetupRoutes(r)
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
