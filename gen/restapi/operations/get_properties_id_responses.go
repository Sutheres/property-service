// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Sutheres/property-service/gen/models"
)

// GetPropertiesIDOKCode is the HTTP code returned for type GetPropertiesIDOK
const GetPropertiesIDOKCode int = 200

/*GetPropertiesIDOK Successful operation

swagger:response getPropertiesIdOK
*/
type GetPropertiesIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Property `json:"body,omitempty"`
}

// NewGetPropertiesIDOK creates GetPropertiesIDOK with default headers values
func NewGetPropertiesIDOK() *GetPropertiesIDOK {

	return &GetPropertiesIDOK{}
}

// WithPayload adds the payload to the get properties Id o k response
func (o *GetPropertiesIDOK) WithPayload(payload *models.Property) *GetPropertiesIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get properties Id o k response
func (o *GetPropertiesIDOK) SetPayload(payload *models.Property) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPropertiesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPropertiesIDCreatedCode is the HTTP code returned for type GetPropertiesIDCreated
const GetPropertiesIDCreatedCode int = 201

/*GetPropertiesIDCreated Created

swagger:response getPropertiesIdCreated
*/
type GetPropertiesIDCreated struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetPropertiesIDCreated creates GetPropertiesIDCreated with default headers values
func NewGetPropertiesIDCreated() *GetPropertiesIDCreated {

	return &GetPropertiesIDCreated{}
}

// WithPayload adds the payload to the get properties Id created response
func (o *GetPropertiesIDCreated) WithPayload(payload interface{}) *GetPropertiesIDCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get properties Id created response
func (o *GetPropertiesIDCreated) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPropertiesIDCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPropertiesIDBadRequestCode is the HTTP code returned for type GetPropertiesIDBadRequest
const GetPropertiesIDBadRequestCode int = 400

/*GetPropertiesIDBadRequest Bad Request

swagger:response getPropertiesIdBadRequest
*/
type GetPropertiesIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetPropertiesIDBadRequest creates GetPropertiesIDBadRequest with default headers values
func NewGetPropertiesIDBadRequest() *GetPropertiesIDBadRequest {

	return &GetPropertiesIDBadRequest{}
}

// WithPayload adds the payload to the get properties Id bad request response
func (o *GetPropertiesIDBadRequest) WithPayload(payload interface{}) *GetPropertiesIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get properties Id bad request response
func (o *GetPropertiesIDBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPropertiesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
