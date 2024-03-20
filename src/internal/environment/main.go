package environment

import "os"

var (
	TOKEN         = os.Getenv("TOKEN")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
	PUBLIC_KEY    = os.Getenv("PUBLIC_KEY")
)
