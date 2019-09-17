// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Sutheres/property-service/gen/models"
)

// GetPropertiesOKCode is the HTTP code returned for type GetPropertiesOK
const GetPropertiesOKCode int = 200

/*GetPropertiesOK list the property addresses

swagger:response getPropertiesOK
*/
type GetPropertiesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Property `json:"body,omitempty"`
}

// NewGetPropertiesOK creates GetPropertiesOK with default headers values
func NewGetPropertiesOK() *GetPropertiesOK {

	return &GetPropertiesOK{}
}

// WithPayload adds the payload to the get properties o k response
func (o *GetPropertiesOK) WithPayload(payload []*models.Property) *GetPropertiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get properties o k response
func (o *GetPropertiesOK) SetPayload(payload []*models.Property) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPropertiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Property, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPropertiesInternalServerErrorCode is the HTTP code returned for type GetPropertiesInternalServerError
const GetPropertiesInternalServerErrorCode int = 500

/*GetPropertiesInternalServerError server error

swagger:response getPropertiesInternalServerError
*/
type GetPropertiesInternalServerError struct {
}

// NewGetPropertiesInternalServerError creates GetPropertiesInternalServerError with default headers values
func NewGetPropertiesInternalServerError() *GetPropertiesInternalServerError {

	return &GetPropertiesInternalServerError{}
}

// WriteResponse to the client
func (o *GetPropertiesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

/*GetPropertiesDefault generic error response

swagger:response getPropertiesDefault
*/
type GetPropertiesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPropertiesDefault creates GetPropertiesDefault with default headers values
func NewGetPropertiesDefault(code int) *GetPropertiesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPropertiesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get properties default response
func (o *GetPropertiesDefault) WithStatusCode(code int) *GetPropertiesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get properties default response
func (o *GetPropertiesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get properties default response
func (o *GetPropertiesDefault) WithPayload(payload *models.Error) *GetPropertiesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get properties default response
func (o *GetPropertiesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPropertiesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
