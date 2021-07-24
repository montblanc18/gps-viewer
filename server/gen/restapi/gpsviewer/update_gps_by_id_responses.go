// Code generated by go-swagger; DO NOT EDIT.

package gpsviewer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/montblanc18/gps-viewer/server/gen/models"
)

// UpdateGpsByIDOKCode is the HTTP code returned for type UpdateGpsByIDOK
const UpdateGpsByIDOKCode int = 200

/*UpdateGpsByIDOK successful operation

swagger:response updateGpsByIdOK
*/
type UpdateGpsByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeviceGPS `json:"body,omitempty"`
}

// NewUpdateGpsByIDOK creates UpdateGpsByIDOK with default headers values
func NewUpdateGpsByIDOK() *UpdateGpsByIDOK {

	return &UpdateGpsByIDOK{}
}

// WithPayload adds the payload to the update gps by Id o k response
func (o *UpdateGpsByIDOK) WithPayload(payload *models.DeviceGPS) *UpdateGpsByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update gps by Id o k response
func (o *UpdateGpsByIDOK) SetPayload(payload *models.DeviceGPS) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGpsByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
