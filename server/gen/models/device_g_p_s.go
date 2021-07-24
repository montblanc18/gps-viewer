// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceGPS device g p s
//
// swagger:model DeviceGPS
type DeviceGPS struct {

	// id
	// Example: 99999
	ID int64 `json:"id,omitempty"`

	// lat
	// Example: 34.69139
	// Required: true
	Lat *float64 `json:"lat"`

	// lng
	// Example: 135.18306
	// Required: true
	Lng *float64 `json:"lng"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *DeviceGPS) UnmarshalJSON(data []byte) error {
	var props struct {

		// id
		// Example: 99999
		ID int64 `json:"id,omitempty"`

		// lat
		// Example: 34.69139
		// Required: true
		Lat *float64 `json:"lat"`

		// lng
		// Example: 135.18306
		// Required: true
		Lng *float64 `json:"lng"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.ID = props.ID
	m.Lat = props.Lat
	m.Lng = props.Lng
	return nil
}

// Validate validates this device g p s
func (m *DeviceGPS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLng(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceGPS) validateLat(formats strfmt.Registry) error {

	if err := validate.Required("lat", "body", m.Lat); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGPS) validateLng(formats strfmt.Registry) error {

	if err := validate.Required("lng", "body", m.Lng); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device g p s based on context it is used
func (m *DeviceGPS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceGPS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceGPS) UnmarshalBinary(b []byte) error {
	var res DeviceGPS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
