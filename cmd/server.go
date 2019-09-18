package cmd

import (
	"github.com/Sutheres/property-service/gen/restapi"
	"github.com/Sutheres/property-service/gen/restapi/operations"
	"github.com/Sutheres/property-service/internal/datastore/database"
	"github.com/Sutheres/property-service/service"
	"github.com/go-openapi/loads"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts Property http server",
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	api := operations.NewPropertyAPI(swaggerSpec)
	if err != nil {
		log.Panicln("Unable to analyze swaggerSpec", err)
	}

	db, err := database.NewDatastore(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicln("database.NewDatastore", err)
	}

	svc := service.NewService(
		BuildDate, CommitHash,
		service.WithDatabase(db),
		)

	service.Configure(api, svc)
	srv := restapi.NewServer(api)
	srv.EnabledListeners = []string{"http"}
	srv.Port = 8080

	defer func() {
		log.Println("shutting down...")
		srv.Shutdown()
	}()

	log.Println("starting server...")

	if err := srv.Serve(); err != nil {
		log.Panicln("server err:", err)
	}
}
