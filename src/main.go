package main

import (
	// standard
	"log/slog"
	"os"

	// external
	"github.com/disgoorg/disgo"
)

var (
	TOKEN         = os.Getenv("TOKEN")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
)

func main() {
	slog.Info("Start")
	slog.Info("DisGo", slog.String("version", disgo.Version))

	disgo.New(
		TOKEN,
	)
}
