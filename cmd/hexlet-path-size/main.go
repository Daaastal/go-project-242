// Package main - точка входа в программу
package main

import (
	"code/internal/app"
	"context"
	"fmt"
	"os"
)

func main() {
	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
