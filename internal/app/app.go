package app

import (
	"code"
	"code/internal/format"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

var errUsage = errors.New("usage: hexlet-path-size [flags] <path>")

func createCmd() *cli.Command {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				DefaultText: "false",
				Usage:       "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				DefaultText: "false",
				Usage:       "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				DefaultText: "false",
				Usage:       "recursive size of directories",
			},
		},
		Action: runMyDuCommand,
	}
	return cmd
}

func runMyDuCommand(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 1 {
		return errUsage
	}

	path := cmd.Args().First()
	human := cmd.Bool("human")
	all := cmd.Bool("all")
	recursive := cmd.Bool("recursive")

	size, err := code.GetPathSize(path, recursive, human, all)
	if err != nil {
		return err
	}

	res := format.GetStringResult(size, path)
	if _, errPrint := fmt.Fprintf(os.Stdout, res); errPrint != nil {
		return errPrint
	}
	return nil
}

func Run(ctx context.Context, osArgs []string) (deferErr error) {
	cmd := createCmd()
	err := cmd.Run(ctx, osArgs)
	return err
}
