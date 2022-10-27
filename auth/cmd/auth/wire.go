package auth

import "os"

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "Auth"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	// prefixs is the config environment variable prefix
	prefixs = []string{"ALPHINIUM_AUTH_", "AUTH_"}

	id, _ = os.Hostname()
)
