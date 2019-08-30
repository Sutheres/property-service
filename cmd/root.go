package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	BuildDate  = "None"
	CommitHash = "None"
)

var rootCmd = &cobra.Command{
	Use: "Property Service",
}

// Execute : runs registered cobra commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
