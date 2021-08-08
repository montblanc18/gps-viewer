// Code generated by go-swagger; DO NOT EDIT.

package gps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/montblanc18/gps-viewer/server/gen/models"
)

// RegisterGpsByDeviceIDCreatedCode is the HTTP code returned for type RegisterGpsByDeviceIDCreated
const RegisterGpsByDeviceIDCreatedCode int = 201

/*RegisterGpsByDeviceIDCreated successful operation

swagger:response registerGpsByDeviceIdCreated
*/
type RegisterGpsByDeviceIDCreated struct {

	/*
	  In: Body
	*/
	Payload *models.DeviceGPS `json:"body,omitempty"`
}

// NewRegisterGpsByDeviceIDCreated creates RegisterGpsByDeviceIDCreated with default headers values
func NewRegisterGpsByDeviceIDCreated() *RegisterGpsByDeviceIDCreated {

	return &RegisterGpsByDeviceIDCreated{}
}

// WithPayload adds the payload to the register gps by device Id created response
func (o *RegisterGpsByDeviceIDCreated) WithPayload(payload *models.DeviceGPS) *RegisterGpsByDeviceIDCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register gps by device Id created response
func (o *RegisterGpsByDeviceIDCreated) SetPayload(payload *models.DeviceGPS) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterGpsByDeviceIDCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterGpsByDeviceIDBadRequestCode is the HTTP code returned for type RegisterGpsByDeviceIDBadRequest
const RegisterGpsByDeviceIDBadRequestCode int = 400

/*RegisterGpsByDeviceIDBadRequest Bad Request

swagger:response registerGpsByDeviceIdBadRequest
*/
type RegisterGpsByDeviceIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterGpsByDeviceIDBadRequest creates RegisterGpsByDeviceIDBadRequest with default headers values
func NewRegisterGpsByDeviceIDBadRequest() *RegisterGpsByDeviceIDBadRequest {

	return &RegisterGpsByDeviceIDBadRequest{}
}

// WithPayload adds the payload to the register gps by device Id bad request response
func (o *RegisterGpsByDeviceIDBadRequest) WithPayload(payload *models.Error) *RegisterGpsByDeviceIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register gps by device Id bad request response
func (o *RegisterGpsByDeviceIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterGpsByDeviceIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterGpsByDeviceIDInternalServerErrorCode is the HTTP code returned for type RegisterGpsByDeviceIDInternalServerError
const RegisterGpsByDeviceIDInternalServerErrorCode int = 500

/*RegisterGpsByDeviceIDInternalServerError Internal Server Error

swagger:response registerGpsByDeviceIdInternalServerError
*/
type RegisterGpsByDeviceIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterGpsByDeviceIDInternalServerError creates RegisterGpsByDeviceIDInternalServerError with default headers values
func NewRegisterGpsByDeviceIDInternalServerError() *RegisterGpsByDeviceIDInternalServerError {

	return &RegisterGpsByDeviceIDInternalServerError{}
}

// WithPayload adds the payload to the register gps by device Id internal server error response
func (o *RegisterGpsByDeviceIDInternalServerError) WithPayload(payload *models.Error) *RegisterGpsByDeviceIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register gps by device Id internal server error response
func (o *RegisterGpsByDeviceIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterGpsByDeviceIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
