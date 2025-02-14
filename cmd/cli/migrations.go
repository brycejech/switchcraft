package cli

import (
	"fmt"
	"switchcraft/core"

	"github.com/spf13/cobra"
)

func registerMigrationsModule(core *core.Core) {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "SwitchCraft CLI migrations module",
	}
	migrateUpCmd(core, migrateCmd)
	migrateDownCmd(core, migrateCmd)

	rootCmd.AddCommand(migrateCmd)
}

func migrateUpCmd(core *core.Core, parentCmd *cobra.Command) {
	upCmd := &cobra.Command{
		Use:   "up",
		Short: "Migrate database all the way up",
		Run: func(_ *cobra.Command, _ []string) {
			if err := core.MigrateUp(); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Successfully ran up migration(s)")
		},
	}

	parentCmd.AddCommand(upCmd)
}

func migrateDownCmd(core *core.Core, parentCmd *cobra.Command) {
	downCmd := &cobra.Command{
		Use:   "down",
		Short: "Migrate database down by a single migration",
		Run: func(_ *cobra.Command, _ []string) {
			if err := core.MigrateDown(); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Successfully ran down migration")
		},
	}

	parentCmd.AddCommand(downCmd)
}
