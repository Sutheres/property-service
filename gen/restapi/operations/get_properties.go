// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPropertiesHandlerFunc turns a function with the right signature into a get properties handler
type GetPropertiesHandlerFunc func(GetPropertiesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPropertiesHandlerFunc) Handle(params GetPropertiesParams) middleware.Responder {
	return fn(params)
}

// GetPropertiesHandler interface for that can handle valid get properties params
type GetPropertiesHandler interface {
	Handle(GetPropertiesParams) middleware.Responder
}

// NewGetProperties creates a new http.Handler for the get properties operation
func NewGetProperties(ctx *middleware.Context, handler GetPropertiesHandler) *GetProperties {
	return &GetProperties{Context: ctx, Handler: handler}
}

/*GetProperties swagger:route GET /properties getProperties

GetProperties get properties API

*/
type GetProperties struct {
	Context *middleware.Context
	Handler GetPropertiesHandler
}

func (o *GetProperties) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPropertiesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
