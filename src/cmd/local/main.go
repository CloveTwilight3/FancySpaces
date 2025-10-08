package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/OliverSchlueter/goutils/middleware"
	"github.com/OliverSchlueter/goutils/sloki"
	"github.com/fancyinnovations/fancyspaces/src/internal/app"
	"github.com/justinas/alice"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logService := sloki.NewService(sloki.Configuration{
		URL:          "http://localhost:3100/loki/api/v1/push",
		Service:      "fancyspaces",
		ConsoleLevel: slog.LevelDebug,
		LokiLevel:    slog.LevelInfo,
		EnableLoki:   false,
		Handlers:     []sloki.LogHandler{},
	})
	slog.SetDefault(slog.New(logService))

	// Connect to databases
	slDB := ConnectSqliteE2E("fancyspaces_e2e.db")

	// Setup HTTP server
	mux := http.NewServeMux()
	port := "8080"

	app.Start(app.Configuration{
		Mux:    mux,
		SQLite: slDB,
	})

	go func() {
		chain := alice.New(
			middleware.RequestLogging,
			middleware.Recovery,
		).Then(mux)

		err := http.ListenAndServe(":"+port, chain)
		if err != nil {
			slog.Error("Could not start server on port "+port, sloki.WrapError(err))
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Started server on http://localhost:%s\n", port))

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	switch <-sig {
	case os.Interrupt:
		slog.Info("Received interrupt signal, shutting down...")

		slog.Info("Shutdown complete")
	}
}

func ConnectSqliteE2E(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		slog.Error("Failed to open sqlite database", sloki.WrapError(err))
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		slog.Error("Failed to ping sqlite database", sloki.WrapError(err))
		os.Exit(1)
	}

	slog.Info("Connected to sqlite database")
	return db
}

func DisconnectSqlite(db *sql.DB) {
	if err := db.Close(); err != nil {
		slog.Error("Failed to close sqlite database", sloki.WrapError(err))
	}

	slog.Info("Disconnected from sqlite database")
}
