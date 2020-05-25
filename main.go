package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/peterbourgon/ff/v3/ffyaml"
	"gitlab.com/ben178/go-starter/pkg/hellocmd"
	"gitlab.com/ben178/go-starter/pkg/logging"
	"gitlab.com/ben178/go-starter/pkg/rootcmd"
)

func main() {
	appName := filepath.Base(os.Args[0])

	fs := flag.NewFlagSet(appName, flag.ContinueOnError)
	var (
		_                       = fs.String("config", "", "config file")
		logHandler              = fs.String("o", "default", "["+logging.GetLogOutputs()+"]")
		logLevel                = fs.String("v", "error", "["+logging.GetLogLevels()+"]")
		rootCommand, rootConfig = rootcmd.New(appName)

		helloCommand = hellocmd.New(rootConfig)
	)

	if err := ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ffyaml.Parser),
		ff.WithEnvVarNoPrefix(),
	); err != nil {
		log.WithError(err).Error("Unable to parse flags")
	}

	rootCommand.Subcommands = []*ffcli.Command{
		helloCommand,
	}

	if err := rootCommand.Parse(os.Args[1:]); err != nil {
		log.WithError(err).Error("Unable to parse commands")
	}

	if err := logging.Configure(logHandler, logLevel); err != nil {
		log.WithError(err).Error("Unable to configure logger")
	}

	if err := rootCommand.Run(context.Background()); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			log.WithError(err).Error("Unable to start")
		}
	}
}
