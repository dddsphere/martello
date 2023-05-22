package main

import (
	"os"

	a "github.com/dddsphere/martello/internal/app"
	l "github.com/dddsphere/martello/internal/log"
)

const (
	name = "user"
	env  = "usr"
)

var (
	log l.Logger = l.NewLogger(l.LogLevel.Info, false)
)

func main() {
	app := a.NewApp(name, env, log)

	err := app.Run()
	if err != nil {
		log.Errorf("%s exit error: %w", err)
		os.Exit(1)
	}
}
