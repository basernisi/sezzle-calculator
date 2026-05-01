package main

import (
	"log"
	"net/http"

	"github.com/basernisi/sezzle-calculator/backend/internal/adapters/auth"
	httpadapter "github.com/basernisi/sezzle-calculator/backend/internal/adapters/http"
	applicationauth "github.com/basernisi/sezzle-calculator/backend/internal/application/auth"
	"github.com/basernisi/sezzle-calculator/backend/internal/application/calculate"
	"github.com/basernisi/sezzle-calculator/backend/internal/domain/calculator"
	"github.com/basernisi/sezzle-calculator/backend/internal/infrastructure/config"
	"github.com/basernisi/sezzle-calculator/backend/internal/infrastructure/logging"
	"github.com/basernisi/sezzle-calculator/backend/internal/infrastructure/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	logger := logging.NewLogger()
	validator, err := auth.NewJWTValidator(cfg.JWTSecret)
	if err != nil {
		log.Fatalf("create jwt validator: %v", err)
	}
	issuer := auth.NewJWTIssuer(cfg.JWTSecret)

	registry := calculate.NewOperationRegistry(
		calculator.AddOperation{},
		calculator.SubtractOperation{},
		calculator.MultiplyOperation{},
		calculator.DivideOperation{},
		calculator.PowerOperation{},
		calculator.SquareRootOperation{},
		calculator.PercentageOperation{},
	)

	service := calculate.NewService(registry)
	handler := httpadapter.NewHandler(service)
	authService := applicationauth.NewService(cfg.DemoClientID, cfg.DemoClientSecret, issuer)
	authHandler := httpadapter.NewAuthHandler(authService)
	router := server.NewRouter(handler, authHandler, validator, logger, cfg.FrontendOrigin)

	logger.Info("starting server", "address", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
