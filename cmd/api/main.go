package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/nataliabu/library-service/internal/database"
	"github.com/nataliabu/library-service/internal/version"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL        string
	httpPort       int
	adminBasicAuth struct {
		username       string
		hashedPassword string
	}
	customerBasicAuth struct {
		username       string
		hashedPassword string
	}
	db struct {
		dsn string
	}
}

type application struct {
	config config
	db     *database.DB
	logger *slog.Logger
	wg     sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	flag.StringVar(&cfg.baseURL, "base-url", "http://localhost:4444", "base URL for the application")
	flag.IntVar(&cfg.httpPort, "http-port", 4444, "port to listen on for HTTP requests")
	flag.StringVar(&cfg.adminBasicAuth.username, "admin-basic-auth-username", "admin", "basic auth username")
	flag.StringVar(&cfg.adminBasicAuth.hashedPassword, "admin-basic-auth-hashed-password", "$2a$12$BoCjlYAkx02ah/BcZzaQ2edgJxM/hg4qKiXh3ig7FuJkpaFTSxhL2", "basic auth password hashed with bcrpyt")
	flag.StringVar(&cfg.customerBasicAuth.username, "customer-basic-auth-username", "customer", "basic auth username")
	flag.StringVar(&cfg.customerBasicAuth.hashedPassword, "customer-basic-auth-hashed-password", "$2a$12$vCQQulqQTOhlanMhxWkMJuEKaXk8dwA6..JHh1iYZrzqHwM0d1mWe", "basic auth password hashed with bcrpyt")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres:postgres@host:5432/db", "postgreSQL DSN")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.db.dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
	}

	return app.serveHTTP()
}
