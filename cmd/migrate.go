package cmd

import (
	"fmt"
	"github.com/Sutheres/property-service/db"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "Run database migrations",
	Run: runMigrations,
}

func runMigrations(cmd *cobra.Command, args []string) {
	migrator, err := db.NewDb()
	if err != nil {
		log.Panicln(errors.Wrap(err, "NewDB"))
	}
	fmt.Println("running migrations...")

	if err := migrator.Migrate(); err != nil {
		log.Panicln(errors.Wrap(err, "Migrate"))
	}
}