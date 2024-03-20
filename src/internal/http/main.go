package http

import (
	// standard
	"context"
	"crypto/ed25519"
	"log/slog"
	"runtime"

	// internal
	"github.com/Malix-off/Discord_GitHub-Utilities/src/internal/environment"
	"github.com/Malix-off/Discord_GitHub-Utilities/src/pkg/errorHandling"

	// external
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	_ "github.com/disgoorg/disgo/discord"
	_ "github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/httpserver"
)

func main() {
	slog.Info("Start")

	slog.Info("Versions",
		slog.String("Go", runtime.Version()),
		slog.String("DisGo", disgo.Version),
	)

	// use custom ed25519 verify implementation
	httpserver.Verify = func(publicKey httpserver.PublicKey, message, sig []byte) bool {
		return ed25519.Verify(publicKey, message, sig)
	}

	client, err := disgo.New(
		environment.TOKEN,

		bot.WithHTTPServerConfigOpts(
			environment.PUBLIC_KEY,

			httpserver.WithURL("/interactions/callback"),
			httpserver.WithAddress(":80"),
		),
		// bot.WithEventListenerFunc(commandListener),
	)
	errorHandling.PanicCheck(err, "Error while creating client")

	err = client.OpenHTTPServer()
	errorHandling.PanicCheck(err, "Error while starting http server")
	defer client.Close(context.TODO())
	slog.Info("Runing…")
}
