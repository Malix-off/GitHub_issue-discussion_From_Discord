package main

import (
	// standard
	"log/slog"
	"os"
	"runtime"

	// external
	"github.com/disgoorg/disgo"
)

var (
	TOKEN         = os.Getenv("TOKEN")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
)

func main() {
	slog.Info("Start")

	slog.Info("Verions",
		slog.String("Go", runtime.Version()),
		slog.String("DisGo", disgo.Version),
	)

	disgo.New(
		TOKEN,
	)
}
