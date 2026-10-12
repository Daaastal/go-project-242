// Package main - точка входа в программу
package main

import (
	"code/internal/app"
	"context"
	"errors"
	"fmt"
	"os"
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
