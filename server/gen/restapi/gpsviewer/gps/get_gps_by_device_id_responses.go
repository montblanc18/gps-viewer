// Code generated by go-swagger; DO NOT EDIT.

package gps

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

// GetGpsByDeviceIDBadRequestCode is the HTTP code returned for type GetGpsByDeviceIDBadRequest
const GetGpsByDeviceIDBadRequestCode int = 400

/*GetGpsByDeviceIDBadRequest Bad Request

swagger:response getGpsByDeviceIdBadRequest
*/
type GetGpsByDeviceIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGpsByDeviceIDBadRequest creates GetGpsByDeviceIDBadRequest with default headers values
func NewGetGpsByDeviceIDBadRequest() *GetGpsByDeviceIDBadRequest {

	return &GetGpsByDeviceIDBadRequest{}
}

// WithPayload adds the payload to the get gps by device Id bad request response
func (o *GetGpsByDeviceIDBadRequest) WithPayload(payload *models.Error) *GetGpsByDeviceIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gps by device Id bad request response
func (o *GetGpsByDeviceIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGpsByDeviceIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGpsByDeviceIDNotFoundCode is the HTTP code returned for type GetGpsByDeviceIDNotFound
const GetGpsByDeviceIDNotFoundCode int = 404

/*GetGpsByDeviceIDNotFound 当該データなし

swagger:response getGpsByDeviceIdNotFound
*/
type GetGpsByDeviceIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGpsByDeviceIDNotFound creates GetGpsByDeviceIDNotFound with default headers values
func NewGetGpsByDeviceIDNotFound() *GetGpsByDeviceIDNotFound {

	return &GetGpsByDeviceIDNotFound{}
}

// WithPayload adds the payload to the get gps by device Id not found response
func (o *GetGpsByDeviceIDNotFound) WithPayload(payload *models.Error) *GetGpsByDeviceIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gps by device Id not found response
func (o *GetGpsByDeviceIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGpsByDeviceIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGpsByDeviceIDInternalServerErrorCode is the HTTP code returned for type GetGpsByDeviceIDInternalServerError
const GetGpsByDeviceIDInternalServerErrorCode int = 500

/*GetGpsByDeviceIDInternalServerError Internal Server Error

swagger:response getGpsByDeviceIdInternalServerError
*/
type GetGpsByDeviceIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGpsByDeviceIDInternalServerError creates GetGpsByDeviceIDInternalServerError with default headers values
func NewGetGpsByDeviceIDInternalServerError() *GetGpsByDeviceIDInternalServerError {

	return &GetGpsByDeviceIDInternalServerError{}
}

// WithPayload adds the payload to the get gps by device Id internal server error response
func (o *GetGpsByDeviceIDInternalServerError) WithPayload(payload *models.Error) *GetGpsByDeviceIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gps by device Id internal server error response
func (o *GetGpsByDeviceIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGpsByDeviceIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
