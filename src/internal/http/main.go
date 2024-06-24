package http

import (
	// Standard
	"context"
	"log/slog"
	"runtime"

	// Custom
	"github.com/Malix-off/Discord_GitHub-Utilities/src/pkg/errorHandling"

	// External
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	_ "github.com/disgoorg/disgo/discord"
	_ "github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/httpserver"

	// Internal
	_ "github.com/Malix-off/Discord_GitHub-Utilities/src/internal/commands"
	"github.com/Malix-off/Discord_GitHub-Utilities/src/internal/environment"
)

func main() {
	slog.Info("Start")

	slog.Info("System",
		slog.Group(
			"OS",

			slog.String("Arch", runtime.GOARCH),
			slog.String("Name", runtime.GOOS),
		),
		slog.Group(
			"Go",

			slog.String("Compiler", runtime.Compiler()),
			slog.String("Version", runtime.Version()),
		),
		slog.Group(
			"Packages",

			slog.String("disgo", disgo.Version),
		),
	)

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
