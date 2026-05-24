package main

import (
	"os"
	"fmt"
	"log"
	"context"
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
				Name:	 "human",
				Aliases: []string{"H"},
				Usage:	 "human-readable sizes (auto-select unit)",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			human := cmd.Bool("human")

			size, err := GetPathSize(path, false, human, false)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s\t%s", size, path)
			return nil
		},
	}

	_ = cmd.Run(context.Background(), os.Args)
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size := int64(0)

	fi, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}

	mode := fi.Mode()
	if mode.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fi, err := os.Lstat(filepath.Join(path, file.Name()))
			if err != nil {
				log.Fatal(err)
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
	newSize := int64(0)
	remains := int64(0)
	count := 0

	if size < 1024 {
		return fmt.Sprintf("%dB", size), nil
	}

	for size != 0 {
		newSize = size
		remains = size % 1024
		size /= 1024
		count++;
	}

	switch count {
	case 1:
		return fmt.Sprintf("%d.%dKB", newSize, remains), nil
	case 2:
		return fmt.Sprintf("%d.%dMB", newSize, remains), nil
	case 3:
		return fmt.Sprintf("%d.%dGB", newSize, remains), nil
	case 4:
		return fmt.Sprintf("%d.%dTB", newSize, remains), nil
	case 5:
		return fmt.Sprintf("%d.%dPB", newSize, remains), nil
	case 6:
		return fmt.Sprintf("%d.%dEB", newSize, remains), nil
	}
	return fmt.Sprintf("%dB", size), nil
}
