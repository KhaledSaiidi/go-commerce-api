package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/KhaledSaiidi/go-commerce/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)
	logger.Info("Connected to DB", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start", "error", err)
		//	log.Printf("Server has failed to start, err: %v", err)
		os.Exit(1)
	}
}
