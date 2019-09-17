package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/solo-service/db"
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
	migrator, err := db.NewDB()
	if err != nil {
		log.Panicln(errors.Wrap(err, "NewDB"))
	}
	fmt.Println("running migrations...")

	if err := migrator.Migrate(); err != nil {
		log.Panicln(errors.Wrap(err, "Migrate"))
	}
}