package main

import (
	"backend/core/services/app"
	"backend/core/services/router"
	"log/slog"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	appService := core_services_app.New()
	routerService := core_services_router.New(appService)

	appService.LoggerService.Info("Starting server...", slog.String("address", appService.ConfigService.HttpServer.Address))

	go appService.WebsocketService.Start()
	go appService.ExchangeWebsocketService.Start()

	server := &http.Server{
		Addr:         appService.ConfigService.HttpServer.Address,
		Handler:      routerService,
		ReadTimeout:  appService.ConfigService.HttpServer.TimeoutRead,
		WriteTimeout: appService.ConfigService.HttpServer.TimeoutWrite,
		IdleTimeout:  appService.ConfigService.HttpServer.TimeoutIdle,
	}

	if err := server.ListenAndServe(); err != nil {
		appService.LoggerService.Error("Failed to start server")
	}

	appService.ExchangeWebsocketService.Stop()

	appService.LoggerService.Info("Server stopped")
}
