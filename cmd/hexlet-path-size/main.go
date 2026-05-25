package main

import (
	"os"
	"fmt"
	"log"
	"context"
	"strings"
	"path/filepath"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:		"hexlet-path-size",
		Usage:		"print size of a file or directory",
		ArgsUsage:	"<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:		"human",
				Aliases:	[]string{"H"},
				DefaultText:	"false",
				Usage:		"human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:		"all",
				Aliases:	[]string{"a"},
				DefaultText:	"false",
				Usage:		"include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:		"recursive",
				Aliases:	[]string{"r"},
				DefaultText:	"false",
				Usage:		"recursive size of directories",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			human := cmd.Bool("human")
			all := cmd.Bool("all")
			recursive := cmd.Bool("recursive")

			size, err := GetPathSize(path, recursive, human, all)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s\t%s", size, path)
			return nil
		},
	}

	_ = cmd.Run(context.Background(), os.Args)
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size := int64(0)

	fi, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if isHidden(fi.Name()) && !all {
		return "", nil
	}

	mode := fi.Mode()
	if mode.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}

		for _, file := range files {
			if isHidden(file.Name()) && !all {
				continue
			}

			fi, err := os.Lstat(filepath.Join(path, file.Name()))
			if err != nil {
				return "", err
			}
			size += fi.Size()
		}
	} else {
		size += fi.Size()
	}

	if human {
		return convertHumanFormat(size)
	}

	return fmt.Sprintf("%dB", size), nil
}

func convertHumanFormat(size int64) (string, error) {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	value := float64(size)
	i := 0

	for value >= 1024 && i < len(units)-1 {
		value /= 1024
		i++
	}

	if i == 0 {
		return fmt.Sprintf("%d%s", size, units[i]), nil
	}

	return fmt.Sprintf("%.1f%s", value, units[i]), nil
}
