package hellocmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"gitlab.com/ben178/go-starter/pkg/rootcmd"
)

// Config for the hello subcommand
type Config struct {
	rootConfig *rootcmd.Config
	greeting   string
}

// New returns a usable ffcli.Command for the hello subcommand.
func New(rootConfig *rootcmd.Config) *ffcli.Command {
	cfg := Config{
		rootConfig: rootConfig,
	}

	fs := flag.NewFlagSet(rootConfig.AppName+" hello", flag.ExitOnError)
	fs.StringVar(&cfg.greeting, "greeting", "Hello", "greeting to use")
	rootConfig.RegisterFlags(fs)

	return &ffcli.Command{
		Name:       "hello",
		ShortUsage: "hello [name]",
		ShortHelp:  "Says hello",
		FlagSet:    fs,
		Exec:       cfg.Exec,
	}
}

// Exec function for this command
func (c *Config) Exec(_ context.Context, args []string) error {
	if len(args) == 0 {
		fmt.Println(c.greeting + " World!")
	} else {
		fmt.Println(c.greeting + " " + args[0] + "!")
	}
	return nil
}
