package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"code"

	"github.com/urfave/cli/v3"
)

var errUsage = errors.New("usage: hexlet-path-size [flags] <path>")

func main() {
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

		Action: func(_ context.Context, cmd *cli.Command) error {
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

			if _, errPrint := fmt.Fprintf(os.Stdout, "%s\t%s\n", size, path); errPrint != nil {
				return errPrint
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		switch {
		case errors.Is(err, errUsage):
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		default:
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
