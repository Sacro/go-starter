package main

import (
	"context"
	"errors"
	"flag"
	"os"

	"github.com/apex/log"
	"github.com/peterbourgon/ff/v3/ffcli"

	"gitlab.com/ben178/go-starter/pkg/logging"
)

var (
	appName     = "main"
	rootFlagSet = flag.NewFlagSet(appName, flag.ExitOnError)
	logLevel    = rootFlagSet.String("v", "error", "["+logging.GetLogLevels()+"]")
	logHandler  = rootFlagSet.String("o", "default", "["+logging.GetLogOutputs()+"]")
)

func main() {
	hello := &ffcli.Command{
		Name:       "hello",
		ShortUsage: "hello",
		ShortHelp:  "Says hello!",
		Exec: func(ctx context.Context, args []string) error {
			log.Info("Hello World!")
			return nil
		},
	}

	root := &ffcli.Command{
		ShortUsage: appName + " [flags] <subcommand>",
		FlagSet:    rootFlagSet,
		Subcommands: []*ffcli.Command{
			hello,
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}

	if err := root.Parse(os.Args[1:]); err != nil {
		log.WithError(err).Error("Unable to parse commands")
	}

	if err := logging.Configure(logHandler, logLevel); err != nil {
		log.WithError(err).Error("Unable to configure logger")
	}

	if err := root.Run(context.Background()); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			log.WithError(err).Error("Unable to start")
		}
	}
}
