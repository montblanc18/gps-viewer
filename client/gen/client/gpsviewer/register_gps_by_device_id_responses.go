// Code generated by go-swagger; DO NOT EDIT.

package gpsviewer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/montblanc18/gps-viewer/client/gen/models"
)

// RegisterGpsByDeviceIDReader is a Reader for the RegisterGpsByDeviceID structure.
type RegisterGpsByDeviceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterGpsByDeviceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRegisterGpsByDeviceIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegisterGpsByDeviceIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegisterGpsByDeviceIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterGpsByDeviceIDCreated creates a RegisterGpsByDeviceIDCreated with default headers values
func NewRegisterGpsByDeviceIDCreated() *RegisterGpsByDeviceIDCreated {
	return &RegisterGpsByDeviceIDCreated{}
}

/* RegisterGpsByDeviceIDCreated describes a response with status code 201, with default header values.

successful operation
*/
type RegisterGpsByDeviceIDCreated struct {
	Payload *models.DeviceGPS
}

func (o *RegisterGpsByDeviceIDCreated) Error() string {
	return fmt.Sprintf("[POST /gqs/{deviceId}][%d] registerGpsByDeviceIdCreated  %+v", 201, o.Payload)
}
func (o *RegisterGpsByDeviceIDCreated) GetPayload() *models.DeviceGPS {
	return o.Payload
}

func (o *RegisterGpsByDeviceIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGPS)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterGpsByDeviceIDBadRequest creates a RegisterGpsByDeviceIDBadRequest with default headers values
func NewRegisterGpsByDeviceIDBadRequest() *RegisterGpsByDeviceIDBadRequest {
	return &RegisterGpsByDeviceIDBadRequest{}
}

/* RegisterGpsByDeviceIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RegisterGpsByDeviceIDBadRequest struct {
	Payload *models.Error
}

func (o *RegisterGpsByDeviceIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /gqs/{deviceId}][%d] registerGpsByDeviceIdBadRequest  %+v", 400, o.Payload)
}
func (o *RegisterGpsByDeviceIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterGpsByDeviceIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterGpsByDeviceIDInternalServerError creates a RegisterGpsByDeviceIDInternalServerError with default headers values
func NewRegisterGpsByDeviceIDInternalServerError() *RegisterGpsByDeviceIDInternalServerError {
	return &RegisterGpsByDeviceIDInternalServerError{}
}

/* RegisterGpsByDeviceIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RegisterGpsByDeviceIDInternalServerError struct {
	Payload *models.Error
}

func (o *RegisterGpsByDeviceIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /gqs/{deviceId}][%d] registerGpsByDeviceIdInternalServerError  %+v", 500, o.Payload)
}
func (o *RegisterGpsByDeviceIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterGpsByDeviceIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
