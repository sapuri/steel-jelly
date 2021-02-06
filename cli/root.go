package cli

import (
	"os"

	"github.com/urfave/cli/v2"
)

func Run() error {
	return NewCmdRoot().Run(os.Args)
}

func NewCmdRoot() *cli.App {
	cmd := cli.NewApp()
	cmd.Version = Version
	cmd.Commands = rootSubCommands()
	return cmd
}

func rootSubCommands() []*cli.Command {
	return []*cli.Command{
		newEroterestCmd(),
	}
}
