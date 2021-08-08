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

// GetGpsByDeviceIDReader is a Reader for the GetGpsByDeviceID structure.
type GetGpsByDeviceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGpsByDeviceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGpsByDeviceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGpsByDeviceIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGpsByDeviceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGpsByDeviceIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGpsByDeviceIDOK creates a GetGpsByDeviceIDOK with default headers values
func NewGetGpsByDeviceIDOK() *GetGpsByDeviceIDOK {
	return &GetGpsByDeviceIDOK{}
}

/* GetGpsByDeviceIDOK describes a response with status code 200, with default header values.

successful operationId
*/
type GetGpsByDeviceIDOK struct {
	Payload *models.DeviceGPS
}

func (o *GetGpsByDeviceIDOK) Error() string {
	return fmt.Sprintf("[GET /gqs/{deviceId}][%d] getGpsByDeviceIdOK  %+v", 200, o.Payload)
}
func (o *GetGpsByDeviceIDOK) GetPayload() *models.DeviceGPS {
	return o.Payload
}

func (o *GetGpsByDeviceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGPS)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGpsByDeviceIDBadRequest creates a GetGpsByDeviceIDBadRequest with default headers values
func NewGetGpsByDeviceIDBadRequest() *GetGpsByDeviceIDBadRequest {
	return &GetGpsByDeviceIDBadRequest{}
}

/* GetGpsByDeviceIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetGpsByDeviceIDBadRequest struct {
	Payload *models.Error
}

func (o *GetGpsByDeviceIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /gqs/{deviceId}][%d] getGpsByDeviceIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetGpsByDeviceIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGpsByDeviceIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGpsByDeviceIDNotFound creates a GetGpsByDeviceIDNotFound with default headers values
func NewGetGpsByDeviceIDNotFound() *GetGpsByDeviceIDNotFound {
	return &GetGpsByDeviceIDNotFound{}
}

/* GetGpsByDeviceIDNotFound describes a response with status code 404, with default header values.

当該データなし
*/
type GetGpsByDeviceIDNotFound struct {
	Payload *models.Error
}

func (o *GetGpsByDeviceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /gqs/{deviceId}][%d] getGpsByDeviceIdNotFound  %+v", 404, o.Payload)
}
func (o *GetGpsByDeviceIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGpsByDeviceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGpsByDeviceIDInternalServerError creates a GetGpsByDeviceIDInternalServerError with default headers values
func NewGetGpsByDeviceIDInternalServerError() *GetGpsByDeviceIDInternalServerError {
	return &GetGpsByDeviceIDInternalServerError{}
}

/* GetGpsByDeviceIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetGpsByDeviceIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetGpsByDeviceIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gqs/{deviceId}][%d] getGpsByDeviceIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGpsByDeviceIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGpsByDeviceIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
