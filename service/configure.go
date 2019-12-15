package service

import (
	"github.com/Sutheres/property-service/gen/models"
	"github.com/Sutheres/property-service/gen/restapi/operations"
	"github.com/Sutheres/property-service/internal/auth"
	"github.com/go-openapi/runtime/middleware"
)

func Configure(api *operations.PropertyAPI, service Service) {

	api.GetPropertiesHandler = operations.GetPropertiesHandlerFunc(
		func(params operations.GetPropertiesParams, a *auth.AuthUser) middleware.Responder {
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
		},
	)


	api.GetPropertiesIDHandler = operations.GetPropertiesIDHandlerFunc(
		func(params operations.GetPropertiesIDParams, a *auth.AuthUser) middleware.Responder {
			property, err := service.GetProperty(params.ID)
			if err != nil {
				return operations.NewGetPropertiesInternalServerError()
			}

			p := property.ToProperty()
			return operations.NewGetPropertiesIDOK().WithPayload(
				&p,
			)
		},
	)


}