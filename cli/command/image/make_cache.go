package image

import (
	"fmt"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

type makeCacheOptions struct {
	terminalOutput string
}

// NewMakeCacheCommand creates a new `docker make-cache` command
func NewMakeCacheCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make-cache [OPTIONS]",
		Short: "Make a new cache",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("You made a new cache!!")
			return nil
		},
	}

	return cmd
}
