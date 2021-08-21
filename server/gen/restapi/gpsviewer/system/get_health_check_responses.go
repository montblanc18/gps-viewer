// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/montblanc18/gps-viewer/server/gen/models"
)

// GetHealthCheckOKCode is the HTTP code returned for type GetHealthCheckOK
const GetHealthCheckOKCode int = 200

/*GetHealthCheckOK Health Chack

swagger:response getHealthCheckOK
*/
type GetHealthCheckOK struct {

	/*
	  In: Body
	*/
	Payload *models.OK `json:"body,omitempty"`
}

// NewGetHealthCheckOK creates GetHealthCheckOK with default headers values
func NewGetHealthCheckOK() *GetHealthCheckOK {

	return &GetHealthCheckOK{}
}

// WithPayload adds the payload to the get health check o k response
func (o *GetHealthCheckOK) WithPayload(payload *models.OK) *GetHealthCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health check o k response
func (o *GetHealthCheckOK) SetPayload(payload *models.OK) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHealthCheckInternalServerErrorCode is the HTTP code returned for type GetHealthCheckInternalServerError
const GetHealthCheckInternalServerErrorCode int = 500

/*GetHealthCheckInternalServerError Internal Server Error

swagger:response getHealthCheckInternalServerError
*/
type GetHealthCheckInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHealthCheckInternalServerError creates GetHealthCheckInternalServerError with default headers values
func NewGetHealthCheckInternalServerError() *GetHealthCheckInternalServerError {

	return &GetHealthCheckInternalServerError{}
}

// WithPayload adds the payload to the get health check internal server error response
func (o *GetHealthCheckInternalServerError) WithPayload(payload *models.Error) *GetHealthCheckInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health check internal server error response
func (o *GetHealthCheckInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthCheckInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}