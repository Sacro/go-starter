package rootcmd

import (
	"context"
	"flag"
	"os"

	"github.com/apex/log"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/peterbourgon/ff/v3/ffyaml"
	"gitlab.com/ben178/go-starter/pkg/logging"
)

// Config for the root command, including flags and types that should be
// available to each subcommand.
type Config struct {
	AppName    string
	Config     string
	LogHandler string
	LogLevel   string
}

// New constructs a usable ffcli.Command and an empty Config. The config's token
// and verbose fields will be set after a successful parse. The caller must
// initialize the config's object API client field.
func New(appName string) (*ffcli.Command, *Config) {
	cfg := Config{
		AppName: appName,
	}

	fs := flag.NewFlagSet(appName, flag.ExitOnError)
	cfg.RegisterFlags(fs)

	if err := ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ffyaml.Parser),
		ff.WithEnvVarNoPrefix(),
	); err != nil {
		log.WithError(err).Fatal("Unable to parse flags")
	}

	return &ffcli.Command{
		ShortUsage: appName + " [flags] <subcommand> [flags] [<arg>...]",
		FlagSet:    fs,
		Exec:       cfg.Exec,
	}, &cfg
}

// RegisterFlags registers the flag fields into the provided flag.FlagSet. This
// helper function allows subcommands to register the root flags into their
// flagsets, creating "global" flags that can be passed after any subcommand at
// the commandline.
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.Config, "config", "", "config path (YAML format)")
	fs.StringVar(&c.LogHandler, "o", "", "["+logging.GetLogOutputs()+"]")
	fs.StringVar(&c.LogLevel, "v", "error", "["+logging.GetLogLevels()+"]")
}

// Exec function for this command.
func (c *Config) Exec(context.Context, []string) error {
	// The root command has no meaning, so if it gets executed,
	// display the usage text to the user instead.
	return flag.ErrHelp
}
