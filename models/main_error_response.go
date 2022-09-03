// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MainErrorResponse main error response
//
// swagger:model main.ErrorResponse
type MainErrorResponse struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this main error response
func (m *MainErrorResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this main error response based on context it is used
func (m *MainErrorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MainErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MainErrorResponse) UnmarshalBinary(b []byte) error {
	var res MainErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
