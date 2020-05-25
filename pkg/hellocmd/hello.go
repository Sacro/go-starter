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
}

// New returns a usable ffcli.Command for the hello subcommand.
func New(rootConfig *rootcmd.Config) *ffcli.Command {
	cfg := Config{
		rootConfig: rootConfig,
	}

	fs := flag.NewFlagSet("hello", flag.ExitOnError)
	rootConfig.RegisterFlags(fs)

	return &ffcli.Command{
		Name:       "hello <name>",
		ShortHelp:  "Says hello",
		ShortUsage: "hello",
		Exec:       cfg.Exec,
	}
}

// Exec function for this command
func (c *Config) Exec(_ context.Context, _ []string) error {
	fmt.Println("Hello World!")
	return nil
}
