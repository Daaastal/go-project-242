// Package main - точка входа в программу
package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Daaastal/go-project-242/internal/app"
)

func main() {
	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if errors.Is(err, app.ErrUsage) {
			os.Exit(2)
		}
		os.Exit(1)
	}
}
