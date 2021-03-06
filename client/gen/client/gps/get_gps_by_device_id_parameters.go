// Code generated by go-swagger; DO NOT EDIT.

package gps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetGpsByDeviceIDParams creates a new GetGpsByDeviceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGpsByDeviceIDParams() *GetGpsByDeviceIDParams {
	return &GetGpsByDeviceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGpsByDeviceIDParamsWithTimeout creates a new GetGpsByDeviceIDParams object
// with the ability to set a timeout on a request.
func NewGetGpsByDeviceIDParamsWithTimeout(timeout time.Duration) *GetGpsByDeviceIDParams {
	return &GetGpsByDeviceIDParams{
		timeout: timeout,
	}
}

// NewGetGpsByDeviceIDParamsWithContext creates a new GetGpsByDeviceIDParams object
// with the ability to set a context for a request.
func NewGetGpsByDeviceIDParamsWithContext(ctx context.Context) *GetGpsByDeviceIDParams {
	return &GetGpsByDeviceIDParams{
		Context: ctx,
	}
}

// NewGetGpsByDeviceIDParamsWithHTTPClient creates a new GetGpsByDeviceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGpsByDeviceIDParamsWithHTTPClient(client *http.Client) *GetGpsByDeviceIDParams {
	return &GetGpsByDeviceIDParams{
		HTTPClient: client,
	}
}

/* GetGpsByDeviceIDParams contains all the parameters to send to the API endpoint
   for the get gps by device Id operation.

   Typically these are written to a http.Request.
*/
type GetGpsByDeviceIDParams struct {

	// DeviceID.
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gps by device Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGpsByDeviceIDParams) WithDefaults() *GetGpsByDeviceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gps by device Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGpsByDeviceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gps by device Id params
func (o *GetGpsByDeviceIDParams) WithTimeout(timeout time.Duration) *GetGpsByDeviceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gps by device Id params
func (o *GetGpsByDeviceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gps by device Id params
func (o *GetGpsByDeviceIDParams) WithContext(ctx context.Context) *GetGpsByDeviceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gps by device Id params
func (o *GetGpsByDeviceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gps by device Id params
func (o *GetGpsByDeviceIDParams) WithHTTPClient(client *http.Client) *GetGpsByDeviceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gps by device Id params
func (o *GetGpsByDeviceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get gps by device Id params
func (o *GetGpsByDeviceIDParams) WithDeviceID(deviceID string) *GetGpsByDeviceIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get gps by device Id params
func (o *GetGpsByDeviceIDParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGpsByDeviceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
