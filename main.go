package main

import (
	"fmt"
	"os"

	"github.com/loozhengyuan/hikctl/cmd"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if err := cmd.Execute(); err != nil {
		return err
	}
	return nil
}
