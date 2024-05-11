package main

import (
	"log/slog"
	"os"

	"github.com/sportgroup-hq/auth/internal/bootstrap"
	"github.com/sportgroup-hq/auth/internal/config"
	"github.com/sportgroup-hq/common-lib/logger"
)

func main() {
	deps, err := bootstrap.Up()
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(logger.NewLogger(os.Stdout, config.Get().Log.Level)))

	if err = deps.HTTPServer.Start(); err != nil {
		panic(err)
	}
}
