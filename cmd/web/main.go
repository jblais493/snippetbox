package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Structs
type application struct {
	logger *slog.Logger
}

// Routing
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
