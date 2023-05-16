package cmd

import (
	"github.com/spf13/cobra"
)

func serverCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "server",
		Short: "core server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	command.AddCommand(serverStartCommand())

	return command
}

func serverStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start core server",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := runApp(coreModule())
			if err != nil {
				return err
			}

			return nil
		},
	}
}
