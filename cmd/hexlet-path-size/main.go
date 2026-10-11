package main

import (
	"context"
	"os"
	"fmt"

	"code/internal/app"
)

func main() {
	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
