// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Annotations annotations
// swagger:model Annotations
type Annotations struct {

	// the remarks of namespace
	LbBusType string `json:"lb-bus-type,omitempty"`

	// the namespace name
	LbGpuModel string `json:"lb-gpu-model,omitempty"`
}

// Validate validates this annotations
func (m *Annotations) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Annotations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Annotations) UnmarshalBinary(b []byte) error {
	var res Annotations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
