package main

import (
	"os"

	mapp "github.com/dddsphere/martello/internal/app"
	mlog "github.com/dddsphere/martello/internal/log"
)

const (
	name = "martello"
	env  = "mtl"
)

var (
	log mlog.Logger = mlog.NewLogger(mlog.LogLevel.Info, false)
)

func main() {
	app := mapp.NewApp(name, env, log)

	err := app.Run()
	if err != nil {
		log.Errorf("%s exit error: %w", err)
		os.Exit(1)
	}
}
