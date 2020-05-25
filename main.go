package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/peterbourgon/ff/v3/ffcli"
	"gitlab.com/ben178/go-starter/pkg/hellocmd"
	"gitlab.com/ben178/go-starter/pkg/logging"
	"gitlab.com/ben178/go-starter/pkg/rootcmd"
)

func main() {
	appName := filepath.Base(os.Args[0])

	var (
		rootCommand, rootConfig = rootcmd.New(appName)
		helloCommand            = hellocmd.New(rootConfig)
	)

	rootCommand.Subcommands = []*ffcli.Command{
		helloCommand,
	}

	if err := rootCommand.Parse(os.Args[1:]); err != nil {
		log.WithError(err).Fatal("Unable to parse commands")
	}

	if err := logging.Configure(&rootConfig.LogHandler, &rootConfig.LogLevel); err != nil {
		log.WithError(err).Fatal("Unable to configure logger")
	}

	if err := rootCommand.Run(context.Background()); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			log.WithError(err).Fatal("Unable to start")
		}
	}
}
