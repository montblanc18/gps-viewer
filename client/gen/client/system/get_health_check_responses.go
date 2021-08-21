// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/montblanc18/gps-viewer/client/gen/models"
)

// GetHealthCheckReader is a Reader for the GetHealthCheck structure.
type GetHealthCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHealthCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetHealthCheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHealthCheckOK creates a GetHealthCheckOK with default headers values
func NewGetHealthCheckOK() *GetHealthCheckOK {
	return &GetHealthCheckOK{}
}

/* GetHealthCheckOK describes a response with status code 200, with default header values.

Health Chack
*/
type GetHealthCheckOK struct {
	Payload *models.OK
}

func (o *GetHealthCheckOK) Error() string {
	return fmt.Sprintf("[GET /health][%d] getHealthCheckOK  %+v", 200, o.Payload)
}
func (o *GetHealthCheckOK) GetPayload() *models.OK {
	return o.Payload
}

func (o *GetHealthCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthCheckInternalServerError creates a GetHealthCheckInternalServerError with default headers values
func NewGetHealthCheckInternalServerError() *GetHealthCheckInternalServerError {
	return &GetHealthCheckInternalServerError{}
}

/* GetHealthCheckInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetHealthCheckInternalServerError struct {
	Payload *models.Error
}

func (o *GetHealthCheckInternalServerError) Error() string {
	return fmt.Sprintf("[GET /health][%d] getHealthCheckInternalServerError  %+v", 500, o.Payload)
}
func (o *GetHealthCheckInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHealthCheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
