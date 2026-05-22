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
		Action: func(_ context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)

			size, err := GetPathSize(path, false, false, false)
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
	var size int64
	size = 0

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
	return fmt.Sprintf("%dB", size), nil
}
