package service

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/property/gen/models"
	"github.com/property/gen/restapi/operations"
)

func Configure(api *operations.PropertyAPI, service Service) {

	api.GetPropertiesHandler = operations.GetPropertiesHandlerFunc(
		func(params operations.GetPropertiesParams) middleware.Responder {
			address := "1000 S Grand Ave, Los Angeles, CA 90015"
			property := &models.Property{
				Address: &address,
				ID:      0,
				TurnKey:  true,
			}

			return operations.NewGetPropertiesOK().WithPayload([]*models.Property{
				property,
			})
		})

}