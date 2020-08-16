package cmd

import (
	"github.com/spf13/cobra"
)

var version = "latest"

var cmdRoot = &cobra.Command{
	Use:   "hikctl",
	Short: "hikctl is a configuration management tool for Hikvision devices",
}

// Execute executes the root command
func Execute() error {
	return cmdRoot.Execute()
}
