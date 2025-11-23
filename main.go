package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	weatherPoint "weather-aggregator/weather/handlers"
)

func main() {
	// simple logger to stdout
	logHandlerOpts := slog.HandlerOptions{}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &logHandlerOpts))
	ctx := context.Background()

	// file logger for dedicated large capture retention
	os.MkdirAll("log", 0755)
	f, err := os.OpenFile(
		fmt.Sprintf("log/%s-weather-log.json",
			time.Now().Format("01-02")),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		logger.Error("error creating log file", "error", err)
	}
	flog := slog.New(slog.NewJSONHandler(f, &logHandlerOpts))
	defer f.Close()

	// define servers
	weatherServerPort := "8080"
	weatherServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", weatherServerPort),
		Handler: weatherPoint.WeatherServerHandler(ctx, logger, flog),
	}
	// TODO metrics server, admin server, etc

	var wg sync.WaitGroup

	// start weather server
	logger.Info("starting weather server on port", "port", weatherServerPort)
	wg.Go(func() {
		if err := weatherServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("weather server error", "error", err)
		}
	})
	// TODO start any additional servers

	// Wait for shutdown signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	logger.Info("shutdown signal received")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := weatherServer.Shutdown(ctx); err != nil {
		logger.Error("weather server shutdown error", "error", err)
	}

	wg.Wait()
	logger.Info("All servers stopped gracefully")
}
