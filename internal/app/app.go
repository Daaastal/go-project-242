// Package app - API запуска и настройки команды
package app

import (
	"context"
	"errors"
	"fmt"

	code "github.com/Daaastal/go-project-242"
	"github.com/Daaastal/go-project-242/internal/format"

	"github.com/urfave/cli/v3"
)

// ErrUsage - ошибка использования команды
var ErrUsage = errors.New("usage: hexlet-path-size [flags] <path>")

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
		Action: run,
	}

	return cmd
}

func run(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 1 {
		return ErrUsage
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
	if _, errPrint := fmt.Fprint(cmd.Writer, res); errPrint != nil {
		return errPrint
	}

	return nil
}

// Run - настроить и запустить команду
func Run(ctx context.Context, osArgs []string) (deferErr error) {
	return createCmd().Run(ctx, osArgs)
}
