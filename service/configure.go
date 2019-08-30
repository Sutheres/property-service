package service

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/property/gen/models"
	m "github.com/property/internal/models"
	"github.com/property/gen/restapi/operations"
	"log"
	_ "github.com/lib/pq"
)

func Configure(api *operations.PropertyAPI, service Service) {

	api.GetPropertiesHandler = operations.GetPropertiesHandlerFunc(
		func(params operations.GetPropertiesParams) middleware.Responder {

			db, err := sqlx.Connect(
				"postgres", "host=localhost port=5432 user=postgres dbname=property password='password' sslmode=disable",
			)
			if err != nil {
				log.Panicln("Connect", err)
			}
			properties := []m.Property{}
			db.Select(&properties, "SELECT * FROM property")


			/*address := "1000 S Grand Ave, Los Angeles, CA 90015"
			property := &models.Property{
				Address: &address,
				ID:      0,
				TurnKey:  true,
			}*/

			var p []*models.Property
			for _, props := range properties {
				p2 := props.ToProperty()
				p = append(p, &p2)
			}
			return operations.NewGetPropertiesOK().WithPayload(p)
		})

}