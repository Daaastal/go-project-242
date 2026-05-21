package main

import (
	"os"
	"context"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:		"hexlet-path-size",
		Usage:		"print size of a file or directory",
		ArgsUsage:	"<path>",
	}

	os.Args = []string{"hexlet-path-size"}

	_ = cmd.Run(context.Background(), os.Args)
}
