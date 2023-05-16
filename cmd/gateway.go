package cmd

import (
	"github.com/spf13/cobra"
)

func gatewayCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "gateway",
		Short: "Manage schema migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	command.AddCommand(gatewayStartCommand())

	return command
}

func gatewayStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Gateway start command",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := runApp(gatewayModule())
			if err != nil {
				return err
			}

			return nil
		},
	}
}
