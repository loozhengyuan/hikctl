package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(cmdVersion)
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of hikctl",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("hikctl %s\n", version)
		return nil
	},
}
