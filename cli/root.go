package cli

import (
	"os"

	"github.com/sapuri/steel-jelly/cli/eroterest"
	"github.com/sapuri/steel-jelly/cli/pornhub"
	"github.com/urfave/cli/v2"
)

func Run() error {
	return NewCmdRoot().Run(os.Args)
}

func NewCmdRoot() *cli.App {
	cmd := cli.NewApp()
	cmd.Usage = "A CLI tool for porn sites"
	cmd.Version = Version
	cmd.Commands = rootSubCommands()
	return cmd
}

func rootSubCommands() []*cli.Command {
	return []*cli.Command{
		eroterest.NewEroterestCmd(),
		pornhub.NewPornhubCmd(),
	}
}
