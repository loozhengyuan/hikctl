package cmd

import (
	"github.com/spf13/cobra"

	"github.com/loozhengyuan/hikctl/internal/pkg/config"
)

func init() {
	cmdRoot.AddCommand(cmdApply)
}

var cmdApply = &cobra.Command{
	Use:   "apply",
	Short: "Applies a configuration to a Hikvision device",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := config.ParseFile(args[0])
		if err != nil {
			return err
		}
		err = conf.Apply()
		if err != nil {
			return err
		}
		return nil
	},
}
