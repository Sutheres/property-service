package service

import (
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/lib/pq"
	"github.com/property/gen/models"
	"github.com/property/gen/restapi/operations"
)

func Configure(api *operations.PropertyAPI, service Service) {

	api.GetPropertiesHandler = operations.GetPropertiesHandlerFunc(
		func(params operations.GetPropertiesParams) middleware.Responder {

			properties, err := service.GetProperties()
			if err != nil {
				return operations.NewGetPropertiesInternalServerError()
			}

			var p []*models.Property
			for _, props := range properties {
				p2 := props.ToProperty()
				p = append(p, &p2)
			}
			return operations.NewGetPropertiesOK().WithPayload(p)
		})

}