// Code generated by go-swagger; DO NOT EDIT.

package gpsviewer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/montblanc18/gps-viewer/server/gen/models"
)

// GetGpsByDeviceIDOKCode is the HTTP code returned for type GetGpsByDeviceIDOK
const GetGpsByDeviceIDOKCode int = 200

/*GetGpsByDeviceIDOK successful operationId

swagger:response getGpsByDeviceIdOK
*/
type GetGpsByDeviceIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeviceGPS `json:"body,omitempty"`
}

// NewGetGpsByDeviceIDOK creates GetGpsByDeviceIDOK with default headers values
func NewGetGpsByDeviceIDOK() *GetGpsByDeviceIDOK {

	return &GetGpsByDeviceIDOK{}
}

// WithPayload adds the payload to the get gps by device Id o k response
func (o *GetGpsByDeviceIDOK) WithPayload(payload *models.DeviceGPS) *GetGpsByDeviceIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gps by device Id o k response
func (o *GetGpsByDeviceIDOK) SetPayload(payload *models.DeviceGPS) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGpsByDeviceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}