package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() error {
	command := &cobra.Command{
		Use:   "core-service",
		Short: "Core Service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	command.AddCommand(dbCommand())
	command.AddCommand(serverCommand())
	command.AddCommand(gatewayCommand())

	return command.Execute()
}
