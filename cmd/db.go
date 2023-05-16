package cmd

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vladazn/go-boilerplate/pkg/db/migration/schema"
	"go.uber.org/fx"
)

func dbCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "db",
		Short: "Manage database",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	command.AddCommand(schemaMigrateCommand())

	return command
}

func schemaMigrateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate schema",
		RunE: func(cmd *cobra.Command, args []string) error {

			var migrator *schema.Migrator

			var err error

			runtimeErrs := make(chan error)

			logrus.SetLevel(logrus.DebugLevel)

			_ = fx.New(
				fx.Populate(&migrator),
				fx.Supply(runtimeErrs),
				schemaMigratorModule(),
				catchError(&err),
			)

			if err != nil {
				return err
			}

			err = migrator.Migrate(context.TODO())
			if err != nil {
				return err
			}

			return nil
		},
	}

	return command
}
