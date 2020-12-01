package image

import (
	"context"
	"fmt"
	"strings"

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
			return runMakeCache(dockerCli, makeCacheOptions{})
		},
	}

	return cmd
}

func runMakeCache(dockerCli command.Cli, options makeCacheOptions) error {
	ctx := context.Background()
	// filler for now, haven't implemented anything in the engine yet
	r := strings.NewReader("{\"Something\": 1}")
	resp, err := dockerCli.Client().ImageMakeCache(ctx, r)
	if resp.Body != nil {
		fmt.Println(resp.Body)
		resp.Body.Close()
	}
	return err
}
