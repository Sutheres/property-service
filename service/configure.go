package service

import (
	"fmt"
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
			home := m.Property{}
			err = db.Get(&home, "SELECT * FROM property WHERE id=$1", 1)
			if err != nil {
				log.Panicln("Get", err)
			}
			fmt.Printf("%#v\n", home)


			/*address := "1000 S Grand Ave, Los Angeles, CA 90015"
			property := &models.Property{
				Address: &address,
				ID:      0,
				TurnKey:  true,
			}*/

			h := home.ToProperty()
			return operations.NewGetPropertiesOK().WithPayload([]*models.Property{
				&h,
			})
		})

}